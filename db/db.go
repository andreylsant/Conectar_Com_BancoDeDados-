package db

import "database/sql"

func ConectaComBanco() *sql.DB {
	conecta := "name dbname password host"
	db, err := sql.Open("postgres", conecta)
	if err != nil {
		panic(err)
	}

	return db
}