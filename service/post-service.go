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
	Delete(ID int64) error
	GetBookByID(ID int64) (*tipos.Post, error)
}

type service struct {
	repositorio repo.PostRepositorio
}

func NewPostService(repo repo.PostRepositorio) PostService {
	return &service{
		repositorio: repo,
	}
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
func (s *service) Criar(post *tipos.Post) (*tipos.Post, error) {
	post.ID = rand.Int63()
	return s.repositorio.Save(post)
}

// Função para encontrar
func (s *service) AcharTodos() ([]tipos.Post, error) {
	return s.repositorio.Encontrados()
}

func (s *service) Delete(ID int64) error {
	return s.repositorio.Delete(ID)
}

func (s *service) GetBookByID(ID int64) (*tipos.Post, error) {
	return s.repositorio.GetBookByID(ID)
}
