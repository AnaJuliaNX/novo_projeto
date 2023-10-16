package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiRota = chi.NewRouter()
)

// Outra forma de fazer rota dessa vez sem o mux, com o chi
func NewChiRouter() Router {
	return &chiRouter{}
}

// Função para o chi lidar com as solicitações que tenha o método GET
func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiRota.Get(uri, f)
}

// Função para o chi lidar com as solicitações HTTP que tenham o método POST
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiRota.Post(uri, f)
}

// Função para o chi lidar com as solicitações HTTP que tenham o método DELETE
func (*chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiRota.Delete(uri, f)
}

// Função para subir o servidor usando o CHI
func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server executando na porta %v", port)
	http.ListenAndServe(port, chiRota)
}
