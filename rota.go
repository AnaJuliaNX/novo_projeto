package main

import (
	"encoding/json"
	"net/http"
)

type DadosPost struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Texto  string `json:"texto"`
}

// Slice of Post
var (
	posts []DadosPost
)

// Dados que vou usar no metodo Post da minha rota
func init() {
	posts = []DadosPost{{Id: 1, Titulo: "Titulo 1", Texto: "Texto 1"}}
}

// Função para ver os dados
func GetPosts(w http.ResponseWriter, r *http.Request) {
	//Cabeçalho
	w.Header().Set("Tipo contido", "aplicação de json")
	//Faço a conversão de struct para json
	resultado, erro := json.Marshal(posts)
	if erro != nil {
		//Mando um statuscode e uma mensagem de erro se houve algum erro
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro ao fazer o marshal do conteúdo"}`))
		return
	}
	//Se foi tudo bem executo esse que vai retornar o meu slice
	w.WriteHeader(http.StatusOK)
	w.Write(resultado)
}

// Função para adicionar mais
func addPos(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body)
}
