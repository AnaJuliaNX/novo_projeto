package service

import (
	"errors"
	"math/rand"

	"github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

type PostService interface {
	Validacao(post *tipos.Post) error            //Para fazer a validação dos dados
	Criar(post *tipos.Post) (*tipos.Post, error) //Para criar dados
	AcharTodos() ([]tipos.Post, error)           //Para buscar dados previamente cadastrados
	Delete(id int) error
}

type service struct{}

var (
	repos repo.PostRepositorio
)

func NewPostService(repo repo.PostRepositorio) PostService {
	return &service{}
}

// Fazendo a validação dos dados digitados
func (*service) Validacao(post *tipos.Post) error {
	//verificando se não foi enviado vazio
	if post == nil {
		erro := errors.New("O post não pode ser vazio")
		return erro
	}
	//Verificando se os campos não estão vazios
	if post.Titulo == "" {
		erro := errors.New("O campo titulo não pode estar vazio")
		return erro
	}
	if post.Autor == "" {
		erro := errors.New("O campo autor não pode estar vazio")
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

func (*service) Delete(id int) error {
	return repo.NewFirestoreRepo().Delete(&tipos.Post{})
}
