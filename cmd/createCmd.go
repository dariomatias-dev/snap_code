package cmd

import (
	"dariomatias-dev/snap_code/cmd/create"

	"github.com/spf13/cobra"
)

var solutionKey string
var solutionFileName string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		create.Create(
			args,
			solutionKey,
			solutionFileName,
		)
	},
}

func init() {
	createCmd.Flags().StringVarP(&solutionKey, "name", "n", "", "Define the key")
	createCmd.Flags().StringVarP(&solutionFileName, "filename", "f", "", "Define the filename")
	rootCmd.AddCommand(createCmd)
}
