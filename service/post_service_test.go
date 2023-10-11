package service

import (
	"testing"

	"github.com/AnaJuliaNX/novo_projeto/tipos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Outra forma de testar as partes do meu código
type MockRepositorio struct {
	mock.Mock
}

func (mock *MockRepositorio) Save(post *tipos.Post) (*tipos.Post, error) {
	argumentos := mock.Called()
	resultado := argumentos.Get(0)
	return resultado.(*tipos.Post), argumentos.Error(1)
}
func (mock *MockRepositorio) Encontrados() ([]tipos.Post, error) {
	argumentos := mock.Called()
	resultado := argumentos.Get(0)
	return resultado.([]tipos.Post), argumentos.Error(1)
}

func TestEncontrados(t *testing.T) {
	mockRepo := new(MockRepositorio)

	post := tipos.Post{ID: 3, Titulo: "1984", Autor: "George Orwell"}
	//Retorno esperado
	mockRepo.On("Encontrados").Return([]tipos.Post{post}, nil)
	testService := NewPostService(mockRepo)
	resultado, _ := testService.AcharTodos()

	//Fazemos a verificação em cada item esperando que sejam compativeis
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 3, resultado[0].ID)
	assert.Equal(t, "1984", resultado[0].Titulo)
	assert.Equal(t, "George Orwell", resultado[0].Autor)
}

// Teste na função para adicionar novos livros
func TestCriar(t *testing.T) {
	MockRepo := new(MockRepositorio)

	post := tipos.Post{ID: 3, Titulo: "1984", Autor: "George Orwell"}
	MockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(MockRepo)
	resultado, erro := testService.Criar(&post)

	//Verifica se os dados são compativeis
	MockRepo.AssertExpectations(t)
	assert.Equal(t, 3, resultado.ID)
	assert.Equal(t, "1984", resultado.Titulo)
	assert.Equal(t, "George Orwell", resultado.Autor)
	assert.Nil(t, erro)
}

// Teste para saber se estpa executando corretamente a validação para nenhum campo encontrado
func TestValidadePost(t *testing.T) {
	//Recebe como valor a minha função NewPostService que tem a função validação nele
	testeService := NewPostService(nil) //talvez passo nil porque nesse caso exige que passe o repo
	erro := testeService.Validacao(nil) //passando um post vazio

	assert.NotNil(t, erro)
	assert.Equal(t, erro.Error(), "O post não pode ser vazio") //Exibe a mensagem de erro
}

// Teste se está executando corretamente a validação para campo de titulo em branco
func TestValidadeTitle(t *testing.T) {
	post := tipos.Post{ID: 3, Titulo: "1984", Autor: "George Orwell"}
	testeService := NewPostService(nil)
	erro := testeService.Validacao(&post)

	assert.NotNil(t, erro)
	assert.Equal(t, erro.Error(), "O campo titulo não pode estar vazio")

}

func TestValidadeAuthor(t *testing.T) {
	post := tipos.Post{ID: 1, Titulo: "1984", Autor: ""}
	testeService := NewPostService(nil)
	erro := testeService.Validacao(&post)

	assert.NotNil(t, erro)
	assert.Equal(t, erro.Error(), "O campo autor não pode estar vazio")
}
