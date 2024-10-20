package database

import (
	"database/sql"
	"log"
	"os"
)

func createTables(
	dbcon *sql.DB,
) {
	content, err := os.ReadFile("cmd\\database\\schemas.sql")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = dbcon.Exec(string(content))
	if err != nil {
		log.Fatalln(err)
	}
}
