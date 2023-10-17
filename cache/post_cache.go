package cache

import "github.com/AnaJuliaNX/novo_projeto/tipos"

type PostCache interface {
	Set(key string, value *tipos.Post) //Associo o identificador do post a um post com uma key especifica
	Get(key string) *tipos.Post        //Recupero o post associado a key especifica
}
