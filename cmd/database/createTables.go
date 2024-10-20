package database

import (
	"dariomatias-dev/snap_code/cmd/utils"
	"database/sql"
	"log"
)

func createTables(
	dbcon *sql.DB,
) {
	filePath := "cmd/database/schemas.sql"

	_, err := dbcon.Exec(utils.ReadFile(filePath))
	if err != nil {
		log.Fatalln(err)
	}
}
