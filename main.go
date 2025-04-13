package main

import (
	"bufio"
	"fmt"
	"mypets/cli"
	"os"
	"strconv"
	"time"
)

const (
	// questionsFile = "cli/questions.txt"
	menuFile = "cli/menu.txt"
)

func main() {
	run()
}

func run() {
	scanner := bufio.NewScanner(os.Stdin)
	// questions := forms.ReadFile(questionsFile)
	menu := cli.ReadFile(menuFile)

	for {
		cli.PrintMenu(menu)
		scanner.Scan()
		action := scanner.Text()
		act, _ := strconv.Atoi(action)

		switch act {
		case 1:
			fmt.Println("Cadastrando pet...")
		case 2:
			fmt.Println("Alterando pet...")
		case 3:
			fmt.Println("Deletando pet...")
		case 4:
			fmt.Println("Listando pets...")
		case 5:
			fmt.Println("Listando pets por critério...")
		case 6:
			return
		default:
			fmt.Println("ERRO: a ação deve ser um número entre 1 e 6")
			time.Sleep(time.Millisecond * 3500)
		}
	}
}
