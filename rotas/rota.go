package rota

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

// Slice of Post
var (
	reposi repo.PostRepositorio = repo.NewFirestoreRepo()
)

// Função para ver os dados
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, erro := repo.NewFirestoreRepo().Encontrados()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro ao obter os dados dos livros"}`))
	}
	//Se foi tudo bem executo esse que vai retornar o meu slice
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// Função para adicionar mais um por vez
func AddPost(w http.ResponseWriter, r *http.Request) {
	//A finalidade disso é porque
	w.Header().Set("content-type", "application/json")
	var post tipos.Post
	//Uso o Decoder quando quero "ler" um valor e escrever esse valor em uma variavel de qualquer tipo
	erro := json.NewDecoder(r.Body).Decode(&post)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro ao fazer o unmarshal"}`))
	}
	post.ID = rand.Int63() //Para gerar ID's aleatórios
	repo.NewFirestoreRepo().Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
