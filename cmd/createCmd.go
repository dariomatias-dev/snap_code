package cmd

import "github.com/spf13/cobra"

var solutionKey string
var solutionFileName string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		Create(args)
	},
}

func init() {
	createCmd.Flags().StringVarP(&solutionKey, "name", "n", "", "Define the key")
	createCmd.Flags().StringVarP(&solutionFileName, "filename", "f", "", "Define the filename")
	rootCmd.AddCommand(createCmd)
}
