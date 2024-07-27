package sliceutils

import strutils "github.com/aleaccurso/tools/strUtils"

// RemoveEndDigits removes trailing digits from each string in a slice
func RemoveEndDigits(slice []string) []string {
	for i, s := range slice {
		slice[i] = strutils.RemoveEndDigits(s)
	}
	return slice
}
