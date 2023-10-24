package repo

import (
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

type PostRepositorio interface {
	Save(post *tipos.Post) (*tipos.Post, error)
	Encontrados() ([]tipos.Post, error)
	Delete(ID int64) error
	GetBookByID(ID int64) (*tipos.Post, error)
}
