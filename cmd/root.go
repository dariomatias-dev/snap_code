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

var isUserCreate bool
var isUserUpdate bool

var setUserNameCmd = &cobra.Command{
	Use:   "username",
	Short: "Set UserName from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			managerUsers := NewManagerUsers()

			if isUserCreate {
				managerUsers.Set(args[0])
			} else if isUserUpdate {
				managerUsers.UpdateByUserName(args[0])
			} else {
				fmt.Println("")
			}
		} else {
			fmt.Println("Enter the GitHub username. Use the flag \"-s\" or \"-u\" to set or update the user.")
		}
	},
}

func init() {
	setUserNameCmd.Flags().BoolVarP(&isUserCreate, "set", "s", false, "Set User")
	setUserNameCmd.Flags().BoolVarP(&isUserUpdate, "update", "u", false, "Update User")
	rootCmd.AddCommand(setUserNameCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
