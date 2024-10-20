package utils

import (
	"log"
	"os"
)

func ReadFile(
	filePath string,
) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	return string(content)
}
