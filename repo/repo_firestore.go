package repo

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/AnaJuliaNX/novo_projeto/tipos"
)

type repo struct{}

func NewFirestoreRepo() PostRepositorio {
	return &repo{}
}

const (
	IdDoProjeto = "novo-projeto"
	NomeColecao = "postados"
)

// Função para adicionar/salvar um novo livro
func (r *repo) Save(post *tipos.Post) (*tipos.Post, error) {
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
func (r *repo) Encontrados() ([]tipos.Post, error) {
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
	//Verificar e arrumar isso depois, ta dizendo que é inutilizavél
	// return posts, nil
}
