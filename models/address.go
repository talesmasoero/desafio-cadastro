package models

import (
	"mypets/cli"
	"strings"
)

type address struct {
	City   string
	Street string
	Number string
}

func setPetAddress(pet *Pet, addr string) {
	city, street, number := SplitAddress(addr)

	// Mesmo sendo contra o princípio da responsabilidade única,
	// foi o melhor lugar que eu e o GPT pensamos para atribuir
	// "NÃO INFORMADO" ao número
	if number == "" {
		number = cli.FormatInput(number)
	}

	pet.Address.City = city
	pet.Address.Street = street
	pet.Address.Number = number
}

// Criar um slice de 3 elementos e usar o copy foi uma
// sacada de gênio para ter, obrigatoriamente, 3 elementos
// sem erro de indexação
// Obs.: essa foi minha e não do GPT

func SplitAddress(address string) (string, string, string) {
	addr := make([]string, 3)
	parts := strings.Split(address, "/")

	copy(addr, parts)
	formatAddress(addr)
	return addr[0], addr[1], addr[2]
}

func formatAddress(addr []string) {
	for i := range addr {
		addr[i] = strings.TrimSpace(addr[i])
	}
}
