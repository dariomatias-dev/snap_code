package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var usernameCmd = &cobra.Command{
	Use:   "username",
	Short: "Set Username from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		managerUsers := NewManagerUsers()

		if len(args) != 0 && args[0] != "" {
			managerUsers.Update(args[0])
		} else {
			fmt.Println("Use `sc username [username]` to set or update the GitHub username.")
		}
	},
}

func init() {
	rootCmd.AddCommand(usernameCmd)
}
