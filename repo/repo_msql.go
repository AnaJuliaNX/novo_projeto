package repo

import (
	"database/sql"
)

// Caso decida usar o banco de dados que eu mesma criei no meu pc aqui está todo o processo de conexão com ele
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
