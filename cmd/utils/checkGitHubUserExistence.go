package utils

import (
	"fmt"
	"log"
	"net/http"
)

func CheckGitHubUserExistence(
	userName string,
) bool {
	resp, err := http.Get(
		fmt.Sprintf(
			"https://api.github.com/users/%s",
			userName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	result := resp.StatusCode != 404

	if !result {
		fmt.Println("User does not exist on GitHub.")
	}

	return result
}
