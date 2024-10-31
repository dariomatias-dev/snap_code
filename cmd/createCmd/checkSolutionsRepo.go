package createcmd

import (
	models "dariomatias-dev/snap_code/cmd/database/models/user"
	"fmt"
	"log"
	"net/http"
)

func checkSolutionsRepo(
	user models.UserModel,
) bool {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/solutions",
		user.Username,
	)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	repoExists := resp.StatusCode != 404

	if !repoExists {
		fmt.Printf(
			"The `solutions` repository does not exist for the user `%s`. Please create it first, then register the solution.\n",
			user.Username,
		)
	}

	return repoExists
}
