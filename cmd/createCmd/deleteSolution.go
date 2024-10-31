package createcmd

import (
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
	"database/sql"
	"fmt"
)

func DeleteSolution(
	dbcon *sql.DB,
	solutionKey string,
) {
	solutionsQueries := solutions.NewSolutionsQueries(dbcon)
	solution := solutionsQueries.GetByKey(solutionKey)

	if solution == nil {
		fmt.Printf("error: the key `%s` does not exist.\n", solutionKey)

		return
	}

	err := solutionsQueries.DeleteByKey(solutionKey)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted key.")
	}
}
