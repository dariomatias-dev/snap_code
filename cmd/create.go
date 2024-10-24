package cmd

import (
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
	"fmt"
	"log"
)

func Create(
	args []string,
) {
	solutionsQueries := solutions.NewSolutionsQueries()

	if len(args) == 0 {
		if solutionKey != "" && solutionFileName != "" {
			err := solutionsQueries.Create(
				solution.SolutionModel{
					Key:      solutionKey,
					FileName: solutionFileName,
				},
			)

			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Command created.")
			}
		} else if solutionKey != "" {
			fmt.Println("Specify the file name using the `-f` flag.")
		} else if solutionFileName != "" {
			fmt.Println("Specify the key name using the `-n` flag.")
		} else {
			fmt.Println("Invalid Command.")
		}

		return
	}

	solution := solutionsQueries.GetByKey(args[0])
	
	if solution == nil {
		fmt.Println("Solution does not exist. Use \"sc create -n [key name] -f [file name]\" to create it.")

		return
	}
}
