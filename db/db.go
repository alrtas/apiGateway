package db

import (
	"database/sql"

	//Fica como empty porque só é necessário somente em runtime
	//Necessário instalação prévia do modulo
	_ "github.com/go-sql-driver/mysql"
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

}
