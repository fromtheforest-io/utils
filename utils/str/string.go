package str // avoids collisions with string package

import (
	"regexp"
	"strings"
	"unicode"
)

// Slugify transforms a string into a lowercase, hyphen-separated slug.
func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace all non-letter/number characters with hyphens
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return '-'
	}, s)

	// Collapse multiple hyphens
	re := regexp.MustCompile(`-+`)
	s = re.ReplaceAllString(s, "-")

	// Trim leading and trailing hyphens
	s = strings.Trim(s, "-")

	return s
}
