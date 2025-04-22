package models

type setters []func(*Pet, string)

func GetSetters() setters {
	return setters{
		setPetName,
		setPetType,
		setPetSex,
		setPetAge,
		setPetWeight,
		setPetBreed,
		setPetAddress,
	}
}

var TypeMap = map[string]string{
	"CACHORRO": "CACHORRO",
	"C":        "CACHORRO",
	"GATO":     "GATO",
	"G":        "GATO",
}

var TypeSex = map[string]string{
	"MACHO": "MACHO",
	"M":     "MACHO",
	"FÊMEA": "FÊMEA",
	"FEMEA": "FÊMEA",
	"F":     "FÊMEA",
}

func normalizeType(typ string) string {
	if normalized, ok := TypeMap[typ]; ok {
		return normalized
	}
	return typ
}

func normalizeSex(sex string) string {
	if normalized, ok := TypeSex[sex]; ok {
		return normalized
	}
	return sex
}
