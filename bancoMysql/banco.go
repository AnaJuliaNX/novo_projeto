package bancomysql

import (
	"database/sql"

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
