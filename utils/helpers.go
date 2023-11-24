package utils

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

var alphaNumeric = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// StrLower returns lowercase s.
func StrLower(s string) string {
	return strings.ToLower(s)
}

// StrTrduim returns s with excess whitespace removed.
func StrTrim(s string) string {
	return strings.TrimSpace(s)
}

// StrRandom returns random string of given length.
func StrRandom(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(result)
}

// panicln takes formatted string and panics with an error.
func panicln(format string, a ...interface{}) {
	log.Panicln(fmt.Errorf(format, a...))
}

func ValidatePhoneNumber(phoneNumber string) bool {
	// Define a regular expression for a phone number starting with "260" followed by 7 digits
	phonePattern := `^260\d{7}$`

	// Compile the regular expression
	regex, err := regexp.Compile(phonePattern)
	if err != nil {
		// Handle error if the regular expression compilation fails
		fmt.Println("Error compiling regex:", err)
		return false
	}

	// Use the regular expression to match the phone number
	return regex.MatchString(phoneNumber)
}
