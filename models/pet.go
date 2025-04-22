package models

type Pet struct {
	Name    string
	Type    string
	Sex     string
	Age     string
	Weight  string
	Breed   string
	Address address
}

func NewPet() *Pet {
	return &Pet{}
}

func setPetName(pet *Pet, name string) {
	pet.Name = name
}

// Utilizar essa função para alguns nomes serem aceitos
// no sexo do pet foi ideia do GPT, não estava conseguindo
// usar o map corretamente

func setPetType(pet *Pet, typ string) {
	pet.Type = normalizeType(typ)
}

func setPetSex(pet *Pet, sex string) {
	pet.Sex = normalizeSex(sex)
}

func setPetAge(pet *Pet, age string) {
	pet.Age = age
}

func setPetWeight(pet *Pet, weight string) {
	pet.Weight = weight
}

func setPetBreed(pet *Pet, breed string) {
	pet.Breed = breed
}
