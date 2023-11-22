package utils

import "regexp"

func ReplacePlaceholders(input string, replacements map[string]string) string {
	regex := regexp.MustCompile(`\[(.*?)\]`)
	result := regex.ReplaceAllStringFunc(input, func(match string) string {
		key := match[1 : len(match)-1] // Removing the square brackets to get the key
		if val, found := replacements[key]; found {
			return val
		}
		return match // If no replacement found, return the original match
	})

	return result
}
