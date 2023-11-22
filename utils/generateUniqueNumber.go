package utils

import (
	"fmt"
	"math/rand"
)

func GenerateNumber(numberType string, code string, date string, total int, separator string) string {

	//t := time.Now()
	//year := t.Year()
	nt := total + 1
	number := fmt.Sprintf("%04d", nt)
	finalNumber := ""

	if separator == "hyphen" {
		finalNumber = fmt.Sprintf("%s-%s-%s-%s", numberType, code, date, number)
	} else {
		finalNumber = fmt.Sprintf("%s/%s/%s/%s", numberType, code, date, number)
	}

	return finalNumber
}

func GenerateRandomDigits(length int) string {
	digits := make([]byte, length)

	for i := 0; i < length; i++ {
		digits[i] = byte(rand.Intn(10) + '0')
	}

	return string(digits)
}
