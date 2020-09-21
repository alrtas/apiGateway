package db

import (
	"database/sql"
	"fmt"

	//Fica como empty porque só é necessário somente em runtime
	//Necessário instalação prévia do modulo
	_ "github.com/go-sql-driver/mysql"
	//Documentação: https://github.com/go-sql-driver/mysql
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis/v8"

	//Documentação: https://github.com/go-redis/redis
	_ "github.com/elastic/go-elasticsearch/v7"
	//Documentação: https://github.com/elastic/go-elasticsearch
)

//ConectaMySQL função que conecta na  base de dados A
func ConectaMySQL() *sql.DB {
	conexao := "root:teste123@s@tcp(127.0.0.1)/baseA"
	db, err := sql.Open("mysql", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

//ConectaElasticSearch função que conecta na base de dados B
func ConectaElasticSearch() {

}

//ConectaRedis função que conecta na base de dados C
func ConectaRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(rdb)
	//return rdb
}
