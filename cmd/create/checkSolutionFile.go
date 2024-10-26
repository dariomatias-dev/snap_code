package create

import (
	"fmt"
	"log"
	"net/http"
)

func checkSolutionFile(
	userName string,
	solutionFileName string,
) bool {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/solutions/contents/%s",
		userName,
		solutionFileName,
	)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	fileExists := resp.StatusCode != 404

	if !fileExists {
		fmt.Printf(
			"The file `%s` does not exist within the `solutions` repository. Ensure the file exists before registering the solution.\n",
			solutionFileName,
		)
	}

	return fileExists
}
