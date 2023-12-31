package repo

import (
	"context"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/AnaJuliaNX/novo_projeto/repo"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

type reposi struct{}

// Comandos de save, show e delete usando o banco firestore
func NewFirestoreRepo() repo.PostRepositorio {
	return &reposi{}
}

const (
	IdDoProjeto = "novo-projeto-3ee53"
	NomeColecao = "postados"
)

// Função para adicionar/salvar um novo livro
func (r *reposi) Save(post *tipos.Post) (*tipos.Post, error) {
	//Para adicionar um novo livro primeiro faz igual o "statement Prepare"
	ctx := context.Background()
	livro, erro := firestore.NewClient(ctx, IdDoProjeto)
	if erro != nil {
		log.Fatalf("Falha ao adicionar um livro no Firestore: %v", erro)
		return nil, erro
	}
	defer livro.Close()

	//Para adicionar os novos dados de um livro no Firestore, como se fosse o "statement Exec"
	_, _, erro = livro.Collection(NomeColecao).Add(ctx, map[string]interface{}{
		"ID":     post.ID,
		"Titulo": post.Titulo,
		"Autor":  post.Autor,
	})
	if erro != nil {
		log.Fatalf("Falha ao acidionar um novo livro: %v", erro)
		return nil, erro
	}

	//Se não tive nenhum erro retorno isso
	return post, nil
}

// Função para exibir todos os livros cadastrados
func (r *reposi) Encontrados() ([]tipos.Post, error) {
	ctx := context.Background()
	livro, erro := firestore.NewClient(ctx, IdDoProjeto)
	if erro != nil {
		log.Fatalf("Erro ao buscar todos os livros cadastrados: %v", erro)
		return nil, erro
	}
	defer livro.Close()

	//Next padrão igual já havia feito antes, lê cada linha e pega os dados delas
	var posts []tipos.Post
	iterador := livro.Collection(NomeColecao).Documents(ctx)
	for {
		documento, erro := iterador.Next()
		if erro != nil {
			log.Fatalf("Erro ao buscar os dados dos livros adicionados: %v", erro)
			return nil, erro
		}
		post := tipos.Post{
			ID:     documento.Data()["ID"].(int64),      //identifica que nesse campo precisa ser um int
			Titulo: documento.Data()["Titulo"].(string), //identifica que nesse campo precisa ser uma string
			Autor:  documento.Data()["Autor"].(string),  //identifica que nesse campo precisa ser uma string
		}
		posts = append(posts, post)
	}
}

// Função para deletar usando o firestore
func (r *reposi) Delete(ID int64) error {
	ctx := context.Background()

	// Inicializando o cliente no firestore
	livro, erro := firestore.NewClient(ctx, IdDoProjeto)
	if erro != nil {
		log.Fatalf("Erro ao criar o cliente Firestore: %v", erro)
		return erro
	}
	defer livro.Close()

	// Criando uma referência ao livro que desejo excluir
	referencia := livro.Collection(NomeColecao).Doc(strconv.Itoa(int(ID)))

	// Deletando o documento
	_, erro = referencia.Delete(ctx)
	if erro != nil {
		log.Fatalf("Falha ao deletar o documento: %v", erro)
		return erro
	}
	//Se deu tudo certo não retorno nada
	return nil
}

// Não fiz uma função para buscar o livro com o banco firestore, ta aqui só pra ele não reclamar
func (r *reposi) GetBookByID(ID int64) (*tipos.Post, error) {
	return &tipos.Post{}, nil
}
