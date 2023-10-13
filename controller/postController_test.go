package controller

import (
	"bytes"
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
	TITULO string = "1984"
	AUTOR  string = "George Orwell"
)

var (
	postRepo       repo.PostRepositorio = repo.NewMysqlRepo()
	postSrv        service.PostService  = service.NewPostService(postRepo)
	postController PostController       = NewPostController(postSrv)
)

func TestAddBook(t *testing.T) {

	//Criar uma nova solicitação HTTP POST
	json := []byte(`{"titulo": "` + TITULO + `","autor": "` + AUTOR + `"}`)
	solicitacao, _ := http.NewRequest("POST", "/livros", bytes.NewBuffer(json))
	//Assign HTTP HandleFunc (controller função AddPost)
	handler := http.HandlerFunc(postController.AddBooks)
	//Registrar a resposta HTTP
	resposta := httptest.NewRecorder() //gravador de resposta
	//"Despachar" a solicitação HTTP
	handler.ServeHTTP(resposta, solicitacao)
	//Adicionar asserções no codigo de status HTTP e na resposta
	status := resposta.Code
	if status != http.StatusOK {
		t.Errorf("retorno incorreto do handler: got %v want %v", status, http.StatusOK)
	}

	//Decode da resposta do HTTP
	var post tipos.Post
	json.NewDecoder(io.Reader(resposta.Body)).Decode(&post) //Erro no decoder que será corrigido logo mais
	//Assert na resposta HTTP
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITULO, post.Titulo)
	assert.Equal(t, AUTOR, post.Autor)

	//Limpar banco de dados, para limpar todo os dados do teste feito
	//Limpo os dados porque quando estou testando acabo criando um novo e coomo não quero ele deleto
	cleanUp(&post)
}

func TestGetBooks(t *testing.T) {

}

// Não criei o delete no mysql nem no firestore então será apenas de exemplo
func cleanUp(post *tipos.Post) {
	postRepo.Delete(post)
}
