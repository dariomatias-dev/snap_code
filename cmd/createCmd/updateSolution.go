package createcmd

import (
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"database/sql"
	"fmt"
)

func UpdateSolution(
	dbcon *sql.DB,
	key string,
	solutionKey string,
	solutionFileName string,
) {
	solutionsQueries := solutions.NewSolutionsQueries(dbcon)

	solutionByKey := solutionsQueries.GetByKey(key)

	if solutionByKey == nil {
		fmt.Printf("error: the key `%s` does not exist.\n", key)

		return
	}

	if solutionFileName != "" {
		usersQueries := users.NewUsersQueries(dbcon)
		user := usersQueries.GetAll()[0]

		if !checkSolutionsRepo(user) {
			return
		}

		if !checkSolutionFile(user.Username, solutionFileName) {
			return
		}
	}

	newSolution := solution.UpdateSolutionModel{}

	if solutionKey != "" {
		newSolution.Key = &solutionKey
	}

	if solutionFileName != "" {
		newSolution.FileName = &solutionFileName
	}

	err := solutionsQueries.UpdateByKey(
		key,
		newSolution,
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated solution")
	}
}
