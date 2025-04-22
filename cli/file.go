package cli

import (
	"log"
	"os"
)

func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
