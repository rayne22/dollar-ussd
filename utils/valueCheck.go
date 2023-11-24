package utils

func ValueIsEqual(array []string, valueToFind string) string {
	for _, v := range array {
		if v == valueToFind {
			return valueToFind
		}
	}
	return ""
}
