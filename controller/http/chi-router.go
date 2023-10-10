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

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiRota.Get(uri, f) //dessa forma que o chi lida com as solicitações HTTP com o método GET
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiRota.Post(uri, f) //dessa forma que o chi lida com as solicitações HTTP com o método POST
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server executando na porta %v", port)
	http.ListenAndServe(port, chiRota)
}
