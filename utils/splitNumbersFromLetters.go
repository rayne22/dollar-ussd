package utils

import (
	"strconv"
	"unicode"
)

// SplitNumbersAndLetters gets numbers from alpha chars
func SplitNumbersAndLetters(s string) (string, int) {
	var letters string
	var numbers int
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters += string(r)
		} else if unicode.IsNumber(r) {
			n, err := strconv.Atoi(string(r))
			if err == nil {
				numbers = numbers*10 + n
			}
		}
	}
	return letters, numbers
}
