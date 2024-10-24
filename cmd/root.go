package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sc",
	Short: "CLI for managing template files",
	Long: `Define and manage template files for quick and easy use in your projects. You can 
copy the desired template files to your project directories with a single command, streamlining your workflow.

Use 'sc [command] --help' for more information about a specific command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			commandRun()
		} else {
			fmt.Println("Welcome to SnapCode! Use 'sc --help' to see the available commands.")
		}
	},
}

var setUserName string
var updateUserUpdate string

var setUserNameCmd = &cobra.Command{
	Use:   "username",
	Short: "Set UserName from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		managerUsers := NewManagerUsers()

		if setUserName != "" {
			managerUsers.Set(setUserName)
		} else if updateUserUpdate != "" {
			managerUsers.UpdateByUserName(setUserName)
		} else {
			fmt.Println("Enter the GitHub username. Use the flag `-s` or `-u` followed by the username to set or update the username.")
		}
	},
}

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
	setUserNameCmd.Flags().StringVarP(&setUserName, "set", "s", "", "Set User")
	setUserNameCmd.Flags().StringVarP(&updateUserUpdate, "update", "u", "", "Update User")
	rootCmd.AddCommand(setUserNameCmd)

	createCmd.Flags().StringVarP(&solutionKey, "name", "n", "", "Define the key")
	createCmd.Flags().StringVarP(&solutionFileName, "filename", "f", "", "Define the filename")
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
