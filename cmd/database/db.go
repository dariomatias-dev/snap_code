package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() *sql.DB {
	dbcon, err := sql.Open("sqlite3", "./dev.db")
	if err != nil {
		log.Fatalln(err)
	}

	createTables(dbcon)

	return dbcon
}
