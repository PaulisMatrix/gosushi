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
	tokens := Tokenize(text)
	// normalise the text
	normalisedTokens := NormaliseFilter(tokens)
	// remove the stopwords
	cleanedTokens := StopWordsFilter(normalisedTokens)
	// stemm the text
	return StemmingFilter(cleanedTokens)

}
