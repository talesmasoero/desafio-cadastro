package service

import (
	"fmt"
	"mypets/cli"
	"mypets/models"
	"mypets/storage"
	"strings"
)

const questionsFile = "cli/forms/questions.txt"

func RegisterPet() {
	pet := models.NewPet()

	questions := strings.Split(cli.ReadFile(questionsFile), "\n")
	setters := models.GetSetters()

	for i, question := range questions {
		set := setters[i]
		input := ""

		fmt.Println(question)
		for {
			input = cli.ReadInput()
			if err := Validate(input, i+1); err != nil {
				fmt.Printf("ERRO: %v\n%s\n", err, question)
				continue
			}
			break
		}

		input = cli.FormatInput(input)
		set(pet, input)
	}
	fmt.Println(pet)
	storage.SaveTXT(pet)
}

// [ ] VER O QUE FAZER COM OS COMMITS PENDENTES
// [ ] CRIAR LÓGICA SEPARADA PARA POPULAR/PREENCHER O SLICE DE SETTERFUNC
//     E PENSAR/VER COM A IA NOMES MELHORES - GETSETTERFUNC, CRIAR ARQUIVOS
//     SEPARADOS PARA OS SETTERS, VOLTAR AO MODEL COM PET E ADDRESS

// [ ] MODULARIZAR MAIS O CÓDIGO, VOLTAR COM A PASTA/PACOTE PET EM MODEL E, TALVEZ,
// EM SERVICE E STORAGE

// [X] REFATORAR O CÓDIGO, ADICIONANDO TOUPPER AO LER QUALQUER INPUT E REMOVENDO
// OS DESNECESSÁRIOS

// DEIXAR PARA DEPOIS:
// [ ] TALVEZ CRIAR UMA STRUCT FIELD COM OS CAMPOS PERGUNTA, SETTER, VALIDATOR E, TALVEZ,
// UM CAMPO QUE IDENTIFICA SE A PERGUNTA É OPICIONAL (LÓGICA DO "NÃO INFORMADO")
