package cmd

import (
	"dariomatias-dev/snap_code/cmd/database"
	models "dariomatias-dev/snap_code/cmd/database/models/user"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"dariomatias-dev/snap_code/cmd/utils"
	"fmt"
)

func NewManagerUsers() ManagerUsers {
	dbcon := database.InitializeDatabase()
	usersQueries := users.NewUsersQueries(dbcon)

	return ManagerUsers{
		usersQueries: usersQueries,
	}
}

type ManagerUsers struct {
	usersQueries *users.UsersQueries
}

func (mu *ManagerUsers) Update(
	username string,
) {
	if mu.usersQueries.Count() == 0 {
		if !utils.CheckGitHubUserExistence(username) {
			return
		}

		mu.usersQueries.Create(
			models.CreateUserModel{
				Username: username,
			},
		)

		fmt.Println("Username set.")
	} else {
		user := mu.usersQueries.GetAll()[0]

		if user.Username == username {
			fmt.Println("The username is already in use.")

			return
		}

		if !utils.CheckGitHubUserExistence(username) {
			return
		}

		mu.usersQueries.UpdateByUsername(
			user.Username,
			models.UpdateUserModel{
				Username: &username,
			},
		)

		fmt.Println(mu.usersQueries.GetAll())

		fmt.Println("Username updated.")
	}
}
