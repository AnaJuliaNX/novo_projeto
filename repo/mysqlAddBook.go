package repo

// import (
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	mysqlBanco "github.com/AnaJuliaNX/novo_projeto/mysqlBanco"
// 	"github.com/AnaJuliaNX/novo_projeto/tipos"
// 	"github.com/gorilla/mux"
// )

// type repo struct{}

// func NewMysqlRepo1() PostRepositorio {
// 	return &repo{}
// }

// // Meu jeito padrão usando mysql, posso tentar fazer alumas mudanças depois pra não ficar gigante
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	corpo, erro := io.ReadAll(r.Body)
// 	if erro != nil {
// 		log.Fatalf("Erro ao ler os dados do corpo: %v", erro)
// 		return
// 	}

// 	var body map[string]interface{}
// 	erro = json.Unmarshal(corpo, &body)
// 	if erro != nil {
// 		log.Fatalf("Erro ao converter para json: %v", erro)
// 		return
// 	}
// 	//Se o titulo e autor não constarem no corpo exibo essa mensagem de erro
// 	if body["titulo"] == nil || body["autor"] == nil {
// 		log.Fatalf("Os campos são obrigatórios")
// 		return
// 	}
// 	//Se o titulo ou autor estiverem vazios mostro essa mensagem de erro
// 	if body["titulo"].(string) == "" || body["autor"].(string) == "" {
// 		log.Fatalf("Os campos são obrigatórios")
// 		return
// 	}

// 	//Conexão com o banco de dados
// 	db, erro := mysqlBanco.ConectarNoBanco()
// 	if erro != nil {
// 		log.Fatalf("Erro ao fazer a conexão com o banco %v", erro)
// 		return
// 	}
// 	livro, erro := db.Prepare("Insert into livros_postadas(titulo, autor) values (?, ?)")
// 	if erro != nil {
// 		log.Fatalf("Erro ao preparar para adicionar livro %v", erro)
// 		return
// 	}
// 	defer livro.Close()

// 	inserir, erro := livro.Exec(body["titulo"], body["autor"])
// 	if erro != nil {
// 		log.Fatalf("Erro aoadicionar o livro: %v", erro)
// 		return
// 	}
// 	_, erro = inserir.LastInsertId()
// 	if erro != nil {
// 		log.Fatalf("Erro ao obter o ID: %v", erro)
// 		return
// 	}
// }

// // Função para selecionar e exibir todos os livros cadastrados, somente o básico
// func GetAllBooks(w http.ResponseWriter, r *http.Request) {
// 	//Abre a conexão com banco de dados
// 	db, erro := mysqlBanco.ConectarNoBanco()
// 	if erro != nil {
// 		log.Fatalf("Erro ao fazer a conexão com o banco de dados: %v", erro)
// 		return
// 	}
// 	defer db.Close()

// 	//O "Query" faz uma consulta nas linhas da tabela buscando pelos dados que pedi (id, titulo, autor)
// 	lines, erro := db.Query("select id,titulo, autor from livros_postadas")
// 	if erro != nil {
// 		log.Fatalf("Erro ao buscar os livros: %v", erro)
// 		return
// 	}
// 	defer lines.Close()

// 	//O "Next" verifica se tem mais linhas para ser escaneada e retorna true or false ou até mesmo um erro
// 	var livros []tipos.Post
// 	for lines.Next() {
// 		var livro tipos.Post
// 		//O "Scan" escaneia a linha atual e busca por todo os dados solicitados
// 		erro := lines.Scan(&livro.ID, &livro.Titulo, &livro.Autor)
// 		if erro != nil {
// 			log.Fatalf("Erro ao escanear livros: %v", erro)
// 			return
// 		}
// 		livros = append(livros, livro)
// 	}
// 	//Tranforma os dados buscados de struct para json
// 	erro = json.NewEncoder(w).Encode(livros)
// 	if erro != nil {
// 		log.Fatalf("Erro ao converter para json: %v", erro)
// 		return
// 	}
// }

// // Buscando um livro pelo ID
// func GetBookByID(w http.ResponseWriter, r *http.Request) {
// 	//Convertendo o parametro da rota de string para int
// 	parametros := mux.Vars(r)
// 	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
// 	if erro != nil {
// 		log.Fatalf("Erro ao converter o parametro para inteiro: %v", erro)
// 		return
// 	}
// 	//Buscando um livro com esse ID
// 	livroencontrado, erro := mysqlBanco.BuscandoUMLivro(int(ID))
// 	if erro != nil {
// 		log.Fatalf("Erro ao buscar o livro pelo ID: %v", erro)
// 		return
// 	}
// 	//Convertendo de struct para json
// 	erro = json.NewEncoder(w).Encode(livroencontrado)
// 	if erro != nil {
// 		log.Fatalf("Erro ao converter para json: %v", erro)
// 		return
// 	}
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	parametros := mux.Vars(r)
// 	//Converto o parametro de string para int
// 	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
// 	if erro != nil {
// 		log.Fatalf("Erro ao converter o parametro para inteiro: %v", erro)
// 		return
// 	}
// 	//Executo o comando que faz a conexão com o banco (mais informações no arquivo "comandosBancoErro")
// 	db, erro := mysqlBanco.ConectarNoBanco()
// 	if erro != nil {
// 		log.Fatalf("Erro ao fazer a conexão com o banco de dados: %v", erro)
// 		return
// 	}
// 	defer db.Close()

// 	//Crio o statement que vai excluir o livro especificado pelo Id
// 	statement, erro := db.Prepare("delete from livros_postadas where id = ?")
// 	if erro != nil {
// 		log.Fatalf("Erro ao criar o statement: %v", erro)
// 		return
// 	}
// 	defer statement.Close()
// 	//Executo o statement e excluo o livro
// 	_, erro = statement.Exec(ID)
// 	if erro != nil {
// 		log.Fatalf("Erro ao executar o statement: %v", erro)
// 		return
// 	}
// }
