package service

import (
	"strings"
	"unicode"
)

func formatFloats(float string) string {
	return strings.Replace(float, ",", ".", 1)
}

func hasValidChars(str string) bool {
	for _, r := range str {
		if !unicode.IsLetter(r) && r != ' ' {
			return false
		}
	}
	return true
}

func isNumeric(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
