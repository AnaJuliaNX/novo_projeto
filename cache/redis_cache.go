package cache

import (
	"encoding/json"
	"time"

	"github.com/AnaJuliaNX/novo_projeto/tipos"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host     string
	db       int
	exipirar time.Duration //Definimos o tempo que o elemento ficará disponivel
}

func NewRedisCache(host string, db int, expi time.Duration) PostCache {
	return &redisCache{
		host:     host, //endereço do servidor Redis que quero estabeler conexão
		db:       db,
		exipirar: expi,
	}
}

// Função para criar um novo cliente redis
func (cache *redisCache) GetClientRedis() *redis.Client {
	//Configurações do banco para criar um cliente
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "", //Sem senha por isso aspas vazias
		DB:       cache.db,
	})

}

// Associando o json post a key
func (cache *redisCache) Set(key string, value *tipos.Post) {
	//Criar novo cliente
	cliente := cache.GetClientRedis()
	//Converto o valor para json e trato o erro
	json, erro := json.Marshal(value)
	if erro != nil {
		panic(erro)
	}

	//O conteúdo do json associado aquela chave tem um tepo para expirar
	cliente.Set(key, json, cache.exipirar*time.Second)
}

func (cache *redisCache) Get(key string) *tipos.Post {
	cliente := cache.GetClientRedis()

	//Passo a key que é o identificador
	valor, erro := cliente.Get(key).Result()
	if erro != nil {
		return nil
	}

	post := tipos.Post{}
	//Faço o unmarshal e direciono o destino pra váriavel post
	erro = json.Unmarshal([]byte(valor), &post)
	if erro != nil {
		panic(erro)
	}
	return &post
}
