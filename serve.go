package main

import (
	"fmt"
	"net/http"

	"github.com/AnaJuliaNX/novo_projeto/cache"
	"github.com/AnaJuliaNX/novo_projeto/controller"
	router "github.com/AnaJuliaNX/novo_projeto/controller/http"
	repo "github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/service"
)

// Desse jeito estaremos:
var (
	postRepositorio repo.PostRepositorio = repo.NewMysqlRepo()
	//Independetes de estruturas
	postService service.PostService = service.NewPostService(postRepositorio)
	//Primeiro valor é a porta, segundo é o banco e o terceiro é quantos segundos ficará disponivel
	postCacheSvr cache.PostCache = cache.NewRedisCache("localhost: 6379", 1, 10)
	//Que é independente de Banco de dados
	postController controller.PostController = controller.NewPostController(postService, postCacheSvr)
	//Que é independete de UI (user interface)
	//Ou seja, posso trocar tanto pra uma chi router quanto pra um mux router
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {

	const port string = ":9000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Funcionando")
	})
	//ROTAS DO CURSO
	httpRouter.GET("/postados", postController.GetAllBooks) //rota para buscar os livros no FIRESTORE
	//Rota para buscar o livro pelo ID com o banco Redis
	//httpRouter.GET("/postados/{id}", postController.GetPostByID)
	httpRouter.POST("/postados", postController.AddBooks) //rota para adicionar os livros

	//ROTAS MYSQL
	httpRouter.POST("/livros", repo.AddBook) //rota para adicionar livros no MYSQL
	httpRouter.GET("/livros", repo.GetAllBooks)
	httpRouter.GET("/livros{id}", repo.GetBookByID)
	httpRouter.DELETE("/livros/{id}", repo.Delete)

	//httpRouter.SERVE(os.Getenv("PORTA"))
	httpRouter.SERVE(port)
}
