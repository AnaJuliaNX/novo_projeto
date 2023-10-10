package repo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	mysqlBanco "github.com/AnaJuliaNX/novo_projeto/mysqlBanco"
)

// Meu jeito padrão usando mysql, posso tentar fazer alumas mudanças depois pra não ficar gigante
func AddBook(w http.ResponseWriter, r *http.Request) {

	corpo, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatalf("Erro ao ler os dados do corpo: %v", erro)
		return
	}

	var body map[string]interface{}
	erro = json.Unmarshal(corpo, &body)
	if erro != nil {
		log.Fatalf("Erro ao converter para json: %v", erro)
		return
	}

	if body["titulo"] == nil || body["autor"] == nil {
		log.Fatalf("Os campos são obrigatórios")
		return
	}
	if body["titulo"].(string) == "" || body["autor"].(string) == "" {
		log.Fatalf("Os campos são obrigatórios")
		return
	}

	db, erro := mysqlBanco.ConectarNoBanco()
	if erro != nil {
		log.Fatalf("Erro ao fazer a conexão com o banco %v", erro)
		return
	}

	livro, erro := db.Prepare("Insert into livros_postadas(titulo, autor) values (?, ?)")
	if erro != nil {
		log.Fatalf("Erro ao preparar para adicionar livro %v", erro)
		return
	}
	defer livro.Close()

	inserir, erro := livro.Exec(body["titulo"], body["autor"])
	if erro != nil {
		log.Fatalf("Erro aoadicionar o livro: %v", erro)
		return
	}

	_, erro = inserir.LastInsertId()
	if erro != nil {
		log.Fatalf("Erro ao obter o ID: %v", erro)
		return
	}
}
