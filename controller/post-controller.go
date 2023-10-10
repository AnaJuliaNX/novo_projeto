package controller

import (
	"encoding/json"
	"net/http"

	service "github.com/AnaJuliaNX/novo_projeto/service"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

var (
	postService service.PostService = service.NewPostService()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, erro := postService.AcharTodos()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Mensagem de erro personalizada
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Erro ao obter os dados dos livros"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post tipos.Post
	erro := json.NewDecoder(r.Body).Decode(&post)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//Mensagem de erro personalizada
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Erro ao fazer o unmarshal"})
		return
	}
	erro1 := postService.Validacao(&post)
	if erro1 != nil {
		//Da um codigo de status do erro
		w.WriteHeader(http.StatusInternalServerError)
		//A mensagem de erro vai depender do erro que der, não foi digitado ou está vazio
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: erro1.Error()})
		return
	}

	resultado, erro2 := postService.Criar(&post)
	if erro2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(tipos.ServiceError{Message: "Erro ao salvar os dados"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
