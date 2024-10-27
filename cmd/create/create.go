package create

import (
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"database/sql"
	"fmt"
	"log"
)

func Create(
	dbcon *sql.DB,
	args []string,
	solutionKey string,
	solutionFileName string,
) {
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

			if !checkSolutionFile(user.Username, solutionFileName) {
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
		fmt.Printf(
			"Solution '%s' does not exist. To create it, use the command: `sc create -n [key name] -f [file name]`.\n",
			args[0],
		)

		return
	}

	createFile(
		args[1],
		user.Username,
		solution.FileName,
	)
}
