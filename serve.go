package main

import (
	"fmt"
	"net/http"

	"github.com/AnaJuliaNX/novo_projeto/controller"
	router "github.com/AnaJuliaNX/novo_projeto/controller/http"
	repo "github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/service"
)

// Desse jeito estaremos:
var (
	postRepositorio repo.PostRepositorio = repo.NewFirestoreRepo()
	//Independetes de estruturas
	postService service.PostService = service.NewPostService(postRepositorio)
	//Que é independente de Banco de dados
	postController controller.PostController = controller.NewPostController(postService)
	//Que é independete de UI (user interface)
	//Ou seja, posso trocar tanto pra uma chi router quanto pra um mux router
	httpRouter router.Router = router.NewMuxRouter()
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
	httpRouter.GET("/livros", repo.GetAllBooks)

	httpRouter.SERVE(port)
}
