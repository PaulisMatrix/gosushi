package textsearch

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// applies the appropriate filters and returns the clean tokens list
func analyze(text string) []string {
	// tokenise the text
	// normalise the text
	// remove the stopwords
	// stemm the text
	return nil
}
