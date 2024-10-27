package create

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
	err := solutionsQueries.DeleteByKey(solutionKey)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted key.")
	}
}
