package rota

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Texto  string `json:"texto"`
}

// Slice of Post
var (
	posts []Post
)

// Dados que vou usar no metodo Post da minha rota
func init() {
	posts = []Post{{Id: 1, Titulo: "Titulo 1", Texto: "Texto 1"}}
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

// Função para adicionar mais um por vez
func AddPos(w http.ResponseWriter, r *http.Request) {
	//Não sei qual está sendoa  finalidade visto que executa perfeitamnete sem
	w.Header().Set("", "application/json")
	var post Post
	//Uso o Decoder quando quero "ler" um valor e escrever esse valor em uma variavel de qualquer tipo
	erro := json.NewDecoder(r.Body).Decode(&post)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro ao fazer o unmarshal}`))
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)

	//w.WriteHeader(http.StatusOK)

	resultado, erro := json.Marshal(post)
	w.Write(resultado)

}
