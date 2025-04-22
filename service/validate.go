package service

import (
	"errors"
	"fmt"
	"mypets/models"
	"strconv"
	"strings"
)

func Validate(input string, index int) error {
	switch index {
	case 1:
		if err := validateName(input); err != nil {
			return fmt.Errorf("nome inválido: %v", err)
		}
	case 2:
		if err := validateType(input); err != nil {
			return fmt.Errorf("tipo inválido: %v", err)
		}
	case 3:
		if err := validateSex(input); err != nil {
			return fmt.Errorf("sexo inválido: %v", err)
		}
	case 4:
		if err := validateAge(input); err != nil {
			return fmt.Errorf("idade inválida: %v", err)
		}
	case 5:
		if err := validateWeight(input); err != nil {
			return fmt.Errorf("peso inválido: %v", err)
		}
	case 6:
		if err := validateBreed(input); err != nil {
			return fmt.Errorf("raça inválida: %v", err)
		}
	case 7:
		if err := validateAddress(input); err != nil {
			return fmt.Errorf("endereço inválido: %v", err)
		}
	}
	return nil
}

// Name validators
func validateName(name string) error {
	if !hasValidName(name) {
		return errors.New("o pet precisa ter NOME e SOBRENOME")
	}
	return nil
}

func hasValidName(name string) bool {
	if !hasValidChars(name) {
		return false
	}
	return strings.Contains(name, " ") && len(name) >= 6
}

// Type validators
func validateType(typ string) error {
	if !hasValidType(typ) {
		return errors.New("o pet precisa ser CACHORRO ou GATO (C/G)")
	}
	return nil
}

func hasValidType(typ string) bool {
	if _, exists := models.TypeMap[typ]; exists {
		return true
	}
	return false
}

// Sex validators
func validateSex(sex string) error {
	if !hasValidSex(sex) {
		return errors.New("o pet deve ser MACHO ou FÊMEA")
	}
	return nil
}

func hasValidSex(sex string) bool {
	if sex == "" {
		return true
	}
	if _, exists := models.TypeSex[sex]; exists {
		return true
	}
	return false
}

// Age validators
func validateAge(age string) error {
	if !hasValidAge(age) {
		return errors.New("o pet deve ter idade de 0 até 20 anos")
	}
	return nil
}

func hasValidAge(age string) bool {
	if age == "" {
		return true
	}
	age = formatFloats(age)

	numAge, err := strconv.ParseFloat(age, 64)
	if err != nil || numAge < 0 || numAge > 20 {
		return false
	}
	return true
}

// Weight validators
func validateWeight(weight string) error {
	if !hasValidWeight(weight) {
		return errors.New("o pet deve ter peso de 0.5 até 60kg")
	}
	return nil
}

func hasValidWeight(weight string) bool {
	if weight == "" {
		return true
	}
	weight = formatFloats(weight)

	numWeight, err := strconv.ParseFloat(weight, 64)
	if err != nil || numWeight < 0.5 || numWeight > 60 {
		return false
	}
	return true
}

// Breed validators
func validateBreed(breed string) error {
	if !hasValidBreed(breed) {
		return errors.New("o pet deve ter uma raça real")
	}
	return nil
}

func hasValidBreed(breed string) bool {
	if breed == "" {
		return true
	}
	if !hasValidChars(breed) {
		return false
	}
	return len(breed) >= 4
}

// Address validators
func validateAddress(address string) error {
	if !hasValidAddress(address) {
		return errors.New("formato esperado: \"CIDADE / RUA / NÚMERO\", número é opcional, duas barras obrigatórias")
	}
	return nil
}

func hasValidAddress(address string) bool {
	if strings.Count(address, "/") != 2 {
		return false
	}
	city, street, number := models.SplitAddress(address)

	if city == "" || street == "" {
		return false
	}
	if !hasValidChars(city) || !hasValidChars(street) {
		return false
	}
	if number != "" && !isNumeric(number) {
		return false
	}
	return true
}
