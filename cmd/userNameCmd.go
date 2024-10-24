package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setUserName string
var updateUserUpdate string

var userNameCmd = &cobra.Command{
	Use:   "username",
	Short: "Set UserName from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		managerUsers := NewManagerUsers()

		if setUserName != "" {
			managerUsers.Set(setUserName)
		} else if updateUserUpdate != "" {
			managerUsers.UpdateByUserName(updateUserUpdate)
		} else {
			fmt.Println("Enter the GitHub username. Use the flag `-s` or `-u` followed by the username to set or update the username.")
		}
	},
}

func init() {
	userNameCmd.Flags().StringVarP(&setUserName, "set", "s", "", "Set User")
	userNameCmd.Flags().StringVarP(&updateUserUpdate, "update", "u", "", "Update User")
	rootCmd.AddCommand(userNameCmd)
}
