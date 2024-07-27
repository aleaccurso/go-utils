package arrays

import (
	"regexp"
	"strings"
)

// RemoveEndDigits removes trailing digits from each string in a slice
func RemoveEndDigits(slice []string) []string {
	re := regexp.MustCompile(`\d+$`)
	for i, s := range slice {
		slice[i] = strings.TrimSpace(re.ReplaceAllString(s, ""))
	}
	return slice
}
