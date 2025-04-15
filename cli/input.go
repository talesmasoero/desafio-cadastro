package cli

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReadInt() int {
	input := ReadInput()
	action, _ := strconv.Atoi(input)
	return action
}
