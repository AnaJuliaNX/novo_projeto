package main

import (
	"fmt"
	"log"
	"net/http"

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
	//Minhas rotas
	router.HandleFunc("/postados", rota.GetPosts).Methods("GET")
	router.HandleFunc("/postados", rota.AddPos).Methods("POST")
	log.Println("Executando na porta", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
