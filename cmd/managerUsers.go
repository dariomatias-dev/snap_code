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

func (mu *ManagerUsers) Set(
	userName string,
) {
	if mu.usersQueries.Count() == 0 {
		if !utils.CheckGitHubUserExistence(userName) {
			return
		}

		mu.usersQueries.Create(
			models.CreateUserModel{
				UserName: userName,
			},
		)

		fmt.Println("Username set.")
	} else {
		fmt.Println("Username already exists. Use `sc username -u [username]` to update the username.")
	}
}

func (mu *ManagerUsers) UpdateByUsername(
	newUserName string,
) {
	if mu.usersQueries.Count() != 0 {
		user := mu.usersQueries.GetAll()[0]

		if user.UserName == newUserName {
			fmt.Println("The username is already in use.")

			return
		}

		if !utils.CheckGitHubUserExistence(newUserName) {
			return
		}

		mu.usersQueries.UpdateByUserName(
			user.UserName,
			models.UpdateUserModel{
				UserName: &newUserName,
			},
		)

		fmt.Println("Username updated.")
	} else {
		fmt.Println("User does not exist.\nUse \"sc username -s [username]\" to set the GitHub username.")
	}
}
