package cmd

import (
	"dariomatias-dev/snap_code/cmd/database"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"fmt"
)

func commandRun() {
	dbcon := database.InitializeDatabase()
	usersQueries := users.NewUsers(dbcon)

	if usersQueries.Count() == 0 {
		fmt.Println("Defina o nome do usuário do GitHub")
	}
}
