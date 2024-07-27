package tools

// RemoveAllEndDigits removes trailing digits from each string in a slice
func RemoveAllEndDigits(slice []string) []string {
	for i, s := range slice {
		slice[i] = RemoveEndDigits(s)
	}
	return slice
}
