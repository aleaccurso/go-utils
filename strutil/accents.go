package strutils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// RemoveAccents removes all accents in the string
func RemoveAccents(s string) string {
	var sb strings.Builder
	for _, r := range norm.NFD.String(s) {
		if unicode.Is(unicode.Mn, r) { // Mn: nonspacing marks
			continue
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

// RemoveEndDigits removes trailing digits from string
func RemoveEndDigits(s string) string {
	re := regexp.MustCompile(`\d+$`)
	return strings.TrimSpace(re.ReplaceAllString(s, ""))
}
