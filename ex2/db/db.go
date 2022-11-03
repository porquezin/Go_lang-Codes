package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDB() *sql.DB {
	conexao := "user=postgres dbname=loja password=test123 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
