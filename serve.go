package main

import (
	"fmt"
	"net/http"

	"github.com/AnaJuliaNX/novo_projeto/controller"
	router "github.com/AnaJuliaNX/novo_projeto/controller/http"
	repo "github.com/AnaJuliaNX/novo_projeto/repo"
)

// Desse jeito estaremos livres da solicitação http pra usar tanto a bibliteca CHI quanto a MUX
var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	//Subindo o servidor
	const port string = ":9000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Funcionando") //Vou exibir aó essa mensagem no postman quando sobre o server
	})
	//ROTAS DO CURSO
	httpRouter.GET("/postados", postController.GetAllBooks) //rota para buscar os livros no FIRESTORE
	httpRouter.POST("/postados", postController.AddBooks)   //rota para adicionar os livros

	//ROTAS MYSQL
	httpRouter.POST("/livros", repo.AddBook) //rota para adicionar livros no MYSQL
	httpRouter.GET("/livros", repo.ShowAllBooks)

	httpRouter.SERVE(port)
}
