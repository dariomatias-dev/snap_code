package cmd

import (
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"database/sql"
	"fmt"
)

func checkUserExists(
	dbcon *sql.DB,
) bool {
	usersQueries := users.NewUsersQueries(dbcon)

	user := usersQueries.GetAll()

	userExists := len(user) != 0

	if !userExists {
		fmt.Println("Set the username. Use `sc username -s [username]` to set the GitHub username.")
	}

	return userExists
}
