package mysql

import (
	"errors"
	"fmt"

	"github.com/AnaJuliaNX/novo_projeto/banco"
	"github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

type repositorio struct {
	db banco.Banco
}

func NewMysqlRepo(db banco.Banco) repo.PostRepositorio {
	return &repositorio{
		db: db,
	}
}

// Função para adicionar/salvar um novo livro
func (r *repositorio) Save(post *tipos.Post) (*tipos.Post, error) {
	db := r.db.ConectarNoBanco()
	defer db.Close()

	livro, erro := db.Prepare("Insert into livros_postadas(titulo, autor) values (?, ?)")
	if erro != nil {
		return &tipos.Post{}, errors.New(fmt.Sprintf("Erro ao preparar para adicionar livro %v", erro))
	}
	defer livro.Close()

	inserir, erro := livro.Exec(post.Titulo, post.Autor)
	if erro != nil {
		return &tipos.Post{}, errors.New(fmt.Sprintf("Erro ao adicionar o livro %v", erro))
	}
	postID, erro := inserir.LastInsertId()
	if erro != nil {
		return &tipos.Post{}, errors.New(fmt.Sprintf("Erro ao obter o ID %v", erro))
	}
	post.ID = postID
	return post, nil
}

// Função para exibir todos os livros cadastrados
func (r *repositorio) Encontrados() ([]tipos.Post, error) {
	// Abre a conexão com banco de dados
	db := r.db.ConectarNoBanco()
	defer db.Close()

	// O "Query" faz uma consulta nas linhas da tabela buscando pelos dados que pedi (id, titulo, autor)
	lines, erro := db.Query("select id,titulo, autor from livros_postadas")
	if erro != nil {
		fmt.Println(erro)
		return []tipos.Post{}, errors.New(fmt.Sprintf("Erro ao obter o ID %v", erro))
	}
	defer lines.Close()

	// O "Next" verifica se tem mais linhas para ser escaneada e retorna true or false ou até mesmo um erro
	var livros []tipos.Post
	for lines.Next() {
		var livro tipos.Post
		//O "Scan" escaneia a linha atual e busca por todo os dados solicitados
		erro := lines.Scan(&livro.ID, &livro.Titulo, &livro.Autor)
		if erro != nil {
			return []tipos.Post{}, errors.New(fmt.Sprintf("Erro ao obter o ID %v", erro))
		}
		livros = append(livros, livro)
	}
	return livros, nil
}

func (r *repositorio) Delete(ID int64) error {
	db := r.db.ConectarNoBanco()
	defer db.Close()
	//Crio o statement que vai excluir o livro especificado pelo Id
	statement, erro := db.Prepare("delete from livros_postadas where id = ?")
	if erro != nil {
		fmt.Println(erro)
		return errors.New(fmt.Sprintf("Erro ao obter o ID %v", erro))
	}
	defer statement.Close()
	//Executo o statement e excluo o livro
	_, erro = statement.Exec(ID)
	if erro != nil {
		fmt.Println(erro)
		return errors.New(fmt.Sprintf("Erro ao obter o ID %v", erro))
	}
	return nil
}

func (r *repositorio) GetBookByID(ID int64) (*tipos.Post, error) {
	//Buscando um livro com esse ID
	db := r.db.ConectarNoBanco()
	defer db.Close()

	linhas, erro := db.Query("select id, titulo, autor from livros_postadas where id = ?", ID)
	if erro != nil {
		fmt.Println(erro)
		return &tipos.Post{}, errors.New("erro ao buscar o livro")
	}
	defer linhas.Close()

	var livro tipos.Post
	if linhas.Next() {
		erro := linhas.Scan(&livro.ID, &livro.Titulo, &livro.Autor)
		if erro != nil {
			fmt.Println(erro)
			return &tipos.Post{}, errors.New("erro ao escanear os dados do livro")
		}
	}
	fmt.Println(livro)
	return &livro, nil
}
