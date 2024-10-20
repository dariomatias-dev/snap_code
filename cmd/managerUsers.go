package cmd

import (
	"dariomatias-dev/snap_code/cmd/database"
	models "dariomatias-dev/snap_code/cmd/database/models/user"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"fmt"
)

func NewManagerUsers() ManagerUsers {
	dbcon := database.InitializeDatabase()
	usersQueries := users.NewUsers(dbcon)

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
	mu.usersQueries.Create(
		models.CreateUserModel{
			UserName: userName,
		},
	)
}

func (mu *ManagerUsers) UpdateByUserName(
	userName string,
) {
	if mu.usersQueries.Count() != 0 {
		mu.usersQueries.UpdateByUserName(
			userName,
			models.UpdateUserModel{
				UserName: &userName,
			},
		)
	} else {
		fmt.Println("User does not exist.\nUse \"sc username [GitHub username] -s\" to set the user.")
	}
}
