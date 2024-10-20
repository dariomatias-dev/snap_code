package cmd

import (
	"dariomatias-dev/snap_code/cmd/database"
)

func commandRun(_ []string) {
	database.InitializeDatabase()
}
