package storage

import (
	"fmt"
	"mypets/models"
	"os"
	"strings"
	"time"
)

const (
	yyyymmdd = "20060102"
	hhmm     = "1504" // 03: 12h clock (am/pm) - 15: 24h clock
)

func SaveTXT(p *models.Pet) {
	fileName := formatFileName(p.Name)
	file, err := os.Create("data/" + fileName)
	if err != nil {
		fmt.Printf("ERRO ao criar o arquivo '%s': %v", fileName, err)
	}
	defer file.Close()

	content := formatFileContent(p)
	file.WriteString(content)
}

func formatFileName(petName string) string {
	timestamp := time.Now().Format(yyyymmdd + "_" + hhmm)
	formatedTime := strings.Replace(timestamp, "_", "T", 1)
	formatedPetName := strings.ReplaceAll(petName, " ", "")

	return fmt.Sprintf("%s-%s.txt", formatedTime, formatedPetName)
}

func formatFileContent(p *models.Pet) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s/%s/%s\n",
		p.Name,
		p.Type,
		p.Sex,
		p.Age,
		p.Weight,
		p.Breed,
		p.Address.City,
		p.Address.Street,
		p.Address.Number,
	)
}
