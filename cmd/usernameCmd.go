package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setUsername string
var updateUsername string

var usernameCmd = &cobra.Command{
	Use:   "username",
	Short: "Set Username from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		managerUsers := NewManagerUsers()

		if setUsername != "" {
			managerUsers.Set(setUsername)
		} else if updateUsername != "" {
			managerUsers.UpdateByUsername(updateUsername)
		} else {
			fmt.Println("Enter the GitHub username. Use the flag `-s` or `-u` followed by the username to set or update the username.")
		}
	},
}

func init() {
	usernameCmd.Flags().StringVarP(&setUsername, "set", "s", "", "Set User")
	usernameCmd.Flags().StringVarP(&updateUsername, "update", "u", "", "Update User")
	rootCmd.AddCommand(usernameCmd)
}
