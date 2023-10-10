package repo

import (
	"encoding/json"
	"log"
	"net/http"

	mysqlBanco "github.com/AnaJuliaNX/novo_projeto/mysqlBanco"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

// Função para selecionar e exibir todos os livros cadastrados, somente o básico
func ShowAllBooks(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados
	db, erro := mysqlBanco.ConectarNoBanco()
	if erro != nil {
		log.Fatalf("Erro ao fazer a conexão com o banco de dados: %v", erro)
		return
	}
	defer db.Close()

	//O "Query" faz uma consulta nas linhas da tabela buscando pelos dados que pedi (id, titulo, autor)
	lines, erro := db.Query("select id,titulo, autor from livros_postadas")
	if erro != nil {
		log.Fatalf("Erro ao buscar os livros: %v", erro)
		return
	}
	defer lines.Close()

	//O "Next" verifica se tem mais linhas para ser escaneada e retorna true or false ou até mesmo um erro
	var livros []tipos.Post
	for lines.Next() {
		var livro tipos.Post
		//O "Scan" escaneia a linha atual e busca por todo os dados solicitados
		erro := lines.Scan(&livro.ID, &livro.Titulo, &livro.Autor)
		if erro != nil {
			log.Fatalf("Erro ao escanear livros: %v", erro)
			return
		}
		livros = append(livros, livro)

	}
	//Tranforma os dados buscados de struct para json
	erro = json.NewEncoder(w).Encode(livros)
	if erro != nil {
		log.Fatalf("Erro ao converter para json: %v", erro)
		return
	}
}
