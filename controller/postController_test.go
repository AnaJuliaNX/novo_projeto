package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AnaJuliaNX/novo_projeto/repo"
	service "github.com/AnaJuliaNX/novo_projeto/service"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
	"github.com/stretchr/testify/assert"
)

const (
	ID     int64  = 123
	TITULO string = "1984"
	AUTOR  string = "George Orwell"
)

var (
	//Posso usar tanto o banco Mysql quanto o Firestore. "NewMysqlRepo"
	postRepo       repo.PostRepositorio = repo.NewFirestoreRepo()
	postSrv        service.PostService  = service.NewPostService(postRepo)
	postController PostController       = NewPostController(postSrv)
)

func TestAddBook(t *testing.T) {
	//Inserir um novo post
	inserir()

	//Criar uma nova solicitação HTTP POST
	json := []byte(`{"titulo": "` + TITULO + `","autor": "` + AUTOR + `"}`)
	solicitacao, _ := http.NewRequest("POST", "/livros", bytes.NewBuffer(json))
	//Atribuindo HTTP HandleFunc (controller função AddPost)
	handler := http.HandlerFunc(postController.AddBooks)
	//Registrar a resposta HTTP
	resposta := httptest.NewRecorder() //gravador de resposta
	//"Despachar" a solicitação HTTP
	handler.ServeHTTP(resposta, solicitacao)
	//Adicionar asserções no código de status HTTP e na resposta
	status := resposta.Code
	if status != http.StatusOK {
		t.Errorf("retorno incorreto do handler: got %v want %v", status, http.StatusOK)
	}
	//Decode da resposta do HTTP
	var post tipos.Post
	//json.NewDecoder(io.Reader(resposta.Body)).Decode(&post)
	//Assert na resposta HTTP
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITULO, post.Titulo)
	assert.Equal(t, AUTOR, post.Autor)

	//Limpo os dados porque quando estou testando acabo criando um novo e como não quero ele deleto
	cleanUp(&post)
}

func TestGetBooks(t *testing.T) {
	//criar uma nova solicitação
	solicitacao, _ := http.NewRequest("GET", "/livros", nil)
	//Atribuindo HTTP HandleFunc (controller função AddPost)
	handler := http.HandlerFunc(postController.GetAllBooks)
	//Registrar a resposta HTTP
	resposta := httptest.NewRecorder() //gravador de resposta
	//"Despachar" a solicitação HTTP
	handler.ServeHTTP(resposta, solicitacao)
	//Adicionar asserções no código de status HTTP e na resposta
	status := resposta.Code
	if status != http.StatusOK {
		t.Errorf("retorno incorreto do handler: got %v want %v", status, http.StatusOK)
	}

	//Decode da resposta do HTTP
	var posts []tipos.Post
	json.NewDecoder(io.Reader(resposta.Body)).Decode(&posts)
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITULO, posts[0].Titulo)

	//cleanUp(&posts[0])
}

// Função para inserir um novo livro no banco de dados
func inserir() {
	var post tipos.Post = tipos.Post{
		ID:     ID,
		Titulo: TITULO,
		Autor:  AUTOR,
	}
	repo.NewFirestoreRepo().Save(&post)
}

// Não criei o delete no mysql nem no firestore então será apenas de exemplo
func cleanUp(post *tipos.Post) {
	postRepo.Delete(post)
}
