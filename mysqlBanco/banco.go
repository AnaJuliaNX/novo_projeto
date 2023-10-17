package bancomysql

import (
	"database/sql"
	"errors"

	"github.com/AnaJuliaNX/novo_projeto/tipos"
	_ "github.com/go-sql-driver/mysql"
)

// Testando se consigo faser funcionar com um banco mysql
func ConectarNoBanco() (*sql.DB, error) {
	stringDeConexao := "func:funcionario@/pragmatica_livraria?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringDeConexao)
	if erro != nil {
		return nil, erro
	}
	erro = db.Ping()
	if erro != nil {
		return nil, erro
	}
	return db, nil
}

func BuscandoUMLivro(ID int) (tipos.Post, error) {
	db, erro := ConectarNoBanco()
	if erro != nil {
		return tipos.Post{}, erro
	}
	defer db.Close()

	linhas, erro := db.Query("select id, titulo, autor from livros_postadas where id = ?", ID)
	if erro != nil {
		return tipos.Post{}, errors.New("erro ao buscar o livro")
	}
	defer linhas.Close()

	var livro tipos.Post
	if linhas.Next() {
		erro := linhas.Scan(&livro.ID, livro.Titulo, &livro.Autor)
		if erro != nil {
			return tipos.Post{}, errors.New("erro ao escanear os dados do livro")
		}
	}
	return livro, nil
}
