package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitializeDatabase() *sql.DB {
	dbcon, err := sql.Open("sqlite3", "./dev.db")
	if err != nil {
		log.Fatalln(err)
	}

	createTables(dbcon)

	return dbcon
}
