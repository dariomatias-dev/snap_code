package cmd

import (
	"dariomatias-dev/snap_code/cmd/create"
	"dariomatias-dev/snap_code/cmd/database"
	"fmt"

	"github.com/spf13/cobra"
)

var solutionKey string
var solutionFileName string
var updateSolution bool
var deleteKey bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		dbcon := database.InitializeDatabase()

		if deleteKey {
			if len(args) != 0 {
				create.DeleteSolution(
					dbcon,
					args[0],
				)
			} else {
				fmt.Println("Enter the key name of the solution you want to delete.")
			}

			return
		} else if updateSolution {
			if len(args) != 0 {
				if solutionKey == "" && solutionFileName == "" {
					fmt.Println("Set the value you want to update with -n for the key name, or -f for the file name.")

					return
				}

				create.UpdateSolution(
					dbcon,
					args[0],
					solutionKey,
					solutionFileName,
				)
			} else {
				fmt.Println("Enter the key name of the solution you want to update.")
			}

			return
		}

		create.Create(
			dbcon,
			args,
			solutionKey,
			solutionFileName,
		)
	},
}

func init() {
	createCmd.Flags().StringVarP(&solutionKey, "name", "n", "", "Define the key")
	createCmd.Flags().StringVarP(&solutionFileName, "filename", "f", "", "Define the filename")
	createCmd.Flags().BoolVarP(&updateSolution, "update", "u", false, "Update solution")
	createCmd.Flags().BoolVarP(&deleteKey, "delete", "d", false, "Delete solution")

	rootCmd.AddCommand(createCmd)
}
