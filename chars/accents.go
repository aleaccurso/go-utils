package chars

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

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
