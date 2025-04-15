package main

import (
	"fmt"
	"mypets/cli"
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
	// questions := forms.ReadFile(questionsFile)
	menu := cli.ReadFile(menuFile)

	for {
		cli.PrintMenu(menu)

		action := cli.ReadInt() // ou ReadAction

		switch action {
		case 1:
			fmt.Println("Cadastrando pet...")
			// pet.Register()
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
