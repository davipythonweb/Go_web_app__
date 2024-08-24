package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// funcao de conexao com db
func Con() *sql.DB {

	conexao := "user=postgres dbname=golang_storage password=Db5$Ades10 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}
