package cmd

import (
	"dariomatias-dev/snap_code/cmd/database"
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/database/queries/solutions"
	"dariomatias-dev/snap_code/cmd/database/queries/users"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Create(
	args []string,
) {
	dbcon := database.InitializeDatabase()

	if !checkUserExists(dbcon) {
		return
	}

	usersQueries := users.NewUsersQueries(dbcon)
	user := usersQueries.GetAll()[0]

	solutionsQueries := solutions.NewSolutionsQueries(dbcon)

	if len(args) == 0 {
		if solutionKey != "" && solutionFileName != "" {
			// Check if the solutions repository exists.
			url := fmt.Sprintf(
				"https://api.github.com/repos/%s/solutions",
				user.UserName,
			)
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			if resp.StatusCode == 404 {
				fmt.Printf(
					"The `solutions` repository does not exist for the user `%s`. Please create it first, then register the solution.\n",
					user.UserName,
				)

				return
			}

			// Check if the specified filename exists within the solutions repository.
			url = fmt.Sprintf(
				"https://api.github.com/repos/%s/solutions/contents/%s",
				user.UserName,
				solutionFileName,
			)
			resp, err = http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			if resp.StatusCode == 404 {
				fmt.Printf(
					"The file `%s` does not exist within the `solutions` repository. Ensure the file exists before registering the solution.\n",
					solutionFileName,
				)

				return
			}

			err = solutionsQueries.Create(
				solution.SolutionModel{
					Key:      solutionKey,
					FileName: solutionFileName,
				},
			)

			if err != nil {
				log.Fatalln(err)
			} else {
				fmt.Println("Command created.")
			}
		} else if solutionKey != "" {
			fmt.Println("Specify the file name using the `-f` flag.")
		} else if solutionFileName != "" {
			fmt.Println("Specify the key name using the `-n` flag.")
		} else {
			fmt.Println("Invalid Command.")
		}

		return
	}

	if len(args) < 2 {
		fmt.Println("The solution key and the destination file path must be specified.")

		return
	}

	solution := solutionsQueries.GetByKey(args[0])

	if solution == nil {
		fmt.Println("Solution does not exist. Use \"sc create -n [key name] -f [file name]\" to create it.")

		return
	}

	path := args[1]

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeAppend)
	} else if !info.IsDir() {
		fmt.Println("The specified path is not a directory.")
	}

	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/solutions/refs/heads/main/%s",
		user.UserName,
		solution.FileName,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(body)
	}

	filePath := fmt.Sprintf("%s/%s", path, solution.FileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(file)
	}
	defer file.Close()

	file.Truncate(0)

	file.Write(body)

	fmt.Printf("File created in `%s`.\n", path)
}
