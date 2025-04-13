package cli

import (
	"fmt"
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

func PrintMenu(menu string) {
	fmt.Printf("\n%s", menu)
}
