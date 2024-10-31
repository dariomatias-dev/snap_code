package cmd

import (
	createcmd "dariomatias-dev/snap_code/cmd/createCmd"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sc",
	Short: "CLI for managing template files",
	Long: `Define and manage template files for quick and easy use in your projects. You can 
copy the desired template files to your project directories with a single command, streamlining your workflow.

Use 'sc [command] --help' for more information about a specific command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to SnapCode! Use 'sc --help' to see the available commands.")
	},
}

func init() {
	createcmd.Load(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
