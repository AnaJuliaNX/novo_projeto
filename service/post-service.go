package service

import (
	"errors"
	"math/rand"

	"github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

type PostService interface {
	Validacao(post *tipos.Post) error
	Criar(post *tipos.Post) (*tipos.Post, error)
	AcharTodos() ([]tipos.Post, error)
}

type service struct{}

func NewPostService() PostService {
	return &service{}
}

// Fazendo a validação dos dados digitados
func (*service) Validacao(post *tipos.Post) error {
	if post == nil {
		erro := errors.New("O post não pode ser vazio")
		return erro
	}
	if post.Titulo == "" || post.Autor == "" {
		erro := errors.New("O campo titulo não pode estar vazio")
		return erro
	}
	return nil
}

// Função para criar
func (*service) Criar(post *tipos.Post) (*tipos.Post, error) {
	post.ID = rand.Int63()
	return repo.NewFirestoreRepo().Save(post)
}

// Função para encontrar
func (*service) AcharTodos() ([]tipos.Post, error) {
	return repo.NewFirestoreRepo().Encontrados()
}
