package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	con := "user=seu_user dbname=n2-web password=sua_senha host=localhost sslmode=disable"
	db, erro := sql.Open("postgres", con)
	if erro != nil {
   		panic(erro.Error())
	}
	return db
}