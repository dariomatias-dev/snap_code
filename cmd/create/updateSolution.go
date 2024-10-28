package create

import (
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
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
