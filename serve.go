package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Subindo o servidor
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Funcionando") //Vou exibir essa mensagem pro usuário
	})
	router.HandleFunc("/postados", GetPosts).Methods("GET")
	log.Println("Executando na porta", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
