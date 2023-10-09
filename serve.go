package main

import (
	"fmt"
	"log"
	"net/http"

	repo "github.com/AnaJuliaNX/novo_projeto/repo"
	rota "github.com/AnaJuliaNX/novo_projeto/rotas"
	"github.com/gorilla/mux"
)

func main() {
	//Subindo o servidor
	router := mux.NewRouter()
	const port string = ":9000"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Funcionando") //Vou exibir essa mensagem pro usuário
	})
	//ROTAS DO CURSO
	router.HandleFunc("/postados", rota.GetPosts).Methods("GET")
	router.HandleFunc("/postados", rota.AddPos).Methods("POST")

	//ROTAS MYSQL
	router.HandleFunc("/livros", repo.AddBook).Methods("POST")

	log.Println("Executando na porta", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
