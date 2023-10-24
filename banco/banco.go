package banco

import "database/sql"

type Banco interface {
	ConectarNoBanco() *sql.DB
}
