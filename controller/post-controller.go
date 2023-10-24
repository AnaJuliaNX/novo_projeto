package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/novo_projeto/cache"
	service "github.com/AnaJuliaNX/novo_projeto/service"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
	"github.com/gorilla/mux"
)

type PostController interface {
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	AddBooks(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.PostService
	cache   cache.PostCache
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	//postService = service
	//PostCache = cache
	return &controller{
		service: service,
		cache:   cache,
	}
}

// Função para buscar os livros cadastrados
func (c *controller) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, erro := c.service.AcharTodos()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Mensagem de erro personalizada
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Erro ao obter os dados dos livros"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// Função para adicionar livros novos
func (c *controller) AddBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post tipos.Post
	erro := json.NewDecoder(r.Body).Decode(&post)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Mensagem de erro personalizada
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Erro ao fazer o unmarshal"})
		return
	}
	erro1 := c.service.Validacao(&post)
	if erro1 != nil {
		//Da um codigo de status do erro
		w.WriteHeader(http.StatusInternalServerError)
		//A mensagem de erro vai depender do erro que der, não foi digitado ou está vazio
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: erro1.Error()})
		return
	}

	resultado, erro2 := c.service.Criar(&post)
	if erro2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Mensagem de erro personalizada
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Erro ao salvar os dados"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

// Função para deletar um livro
func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	parametros := mux.Vars(r)
	bookID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Nenhum livro encontrado"})
		return
	}
	erro = c.service.Delete(bookID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: erro.Error()})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso")
}

// APENAS PARA SABER COMO FUNCIONA E QUE EXISTE A POSSIBILIDADE DE TER
func (c *controller) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	parametros := mux.Vars(r)
	bookID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Nenhum livro encontrado"})
		return
	}

	//Associo os livros a essa chave
	var book *tipos.Post = c.cache.Get(parametros["id"])
	//Caso não encontre nenhum livro associado a essa chave
	if book == nil {
		book, erro := c.service.GetBookByID(bookID)
		if erro != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Nenhum livro encontrado"})
			return
		}
		//Se obtermos um valor do service armazenamos esse valor no cache
		c.cache.Set(parametros["id"], book)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}
