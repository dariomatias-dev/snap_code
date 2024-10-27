package create

import (
	"dariomatias-dev/snap_code/cmd/database"
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"fmt"
	"log"
)

func Create(
	args []string,
	solutionKey string,
	solutionFileName string,
) {
	dbcon := database.InitializeDatabase()

	if !checkUserExists(dbcon) {
		return
	}

	usersQueries := users.NewUsersQueries(dbcon)
	user := usersQueries.GetAll()[0]

	solutionsQueries := solutions.NewSolutionsQueries(dbcon)

	if len(args) == 0 {
		if solutionKey != "" && solutionFileName != "" {
			if !checkSolutionsRepo(user) {
				return
			}

			if !checkSolutionFile(user.UserName, solutionFileName) {
				return
			}

			err := solutionsQueries.Create(
				solution.SolutionModel{
					Key:      solutionKey,
					FileName: solutionFileName,
				},
			)

			if err != nil {
				log.Fatalln(err)
			} else {
				fmt.Println("Command created.")
			}
		} else if solutionKey != "" {
			fmt.Println("Specify the file name using the `-f` flag.")
		} else if solutionFileName != "" {
			fmt.Println("Specify the key name using the `-n` flag.")
		} else {
			fmt.Println("Invalid Command. Use `sc help` to get help.")
		}

		return
	}

	if len(args) < 2 {
		fmt.Println("The solution key and the destination file path must be specified.")

		return
	}

	solution := solutionsQueries.GetByKey(args[0])

	if solution == nil {
		fmt.Println("Solution does not exist. Use \"sc create -n [key name] -f [file name]\" to create it.")

		return
	}

	createFile(
		args[1],
		user.UserName,
		solution.FileName,
	)
}
