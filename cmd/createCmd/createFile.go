package createcmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func createFile(
	path string,
	userName string,
	fileName string,
) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeAppend)
	} else if !info.IsDir() {
		fmt.Println("The specified path is not a directory.")
	}

	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/solutions/refs/heads/main/%s",
		userName,
		fileName,
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

	filePath := fmt.Sprintf("%s/%s", path, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(file)
	}
	defer file.Close()

	file.Truncate(0)

	file.Write(body)

	fmt.Printf("File created in `%s`.\n", path)
}