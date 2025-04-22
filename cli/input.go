package cli

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const notInformed = "NÃO INFORMADO"

var reader = bufio.NewReader(os.Stdin)

func ReadInput() string {
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))
	return input
}

func ReadAction() int {
	input := ReadInput()
	action, _ := strconv.Atoi(input)
	return action
}

// FormatInput set input as "NÃO INFORMADO" if it's equals nil ("")
func FormatInput(input string) string {
	if input == "" {
		return notInformed
	}
	return input
}
