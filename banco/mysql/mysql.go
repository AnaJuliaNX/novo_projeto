package bancomysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct{}

// Conexão com banco mysql
func (*DB) ConectarNoBanco() *sql.DB {
	stringDeConexao := "func:funcionario@/pragmatica_livraria?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringDeConexao)
	if erro != nil {
		log.Fatal("Não foi possivel se conecar com o banco")
		return nil
	}
	erro = db.Ping()
	if erro != nil {
		log.Fatal("Não foi possivel se conecar com o banco")
		return nil
	}
	return db
}
