package textsearch

import (
	"strings"

	"github.com/kljensen/snowball"
)

func getStopWords() map[string]struct{} {
	return map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
}

func NormaliseFilter(tokens []string) []string {
	normalisedTokens := make([]string, len(tokens))
	for i, token := range tokens {
		normalisedTokens[i] = strings.ToLower(token)
	}
	return normalisedTokens
}

func StopWordsFilter(tokens []string) []string {
	stopWordsTokens := make([]string, 0, len(tokens))
	stopWordsMap := getStopWords()

	for _, token := range tokens {
		if _, ok := stopWordsMap[token]; !ok {
			stopWordsTokens = append(stopWordsTokens, token)
		}
	}
	return stopWordsTokens
}

func StemmingFilter(tokens []string) []string {
	stemmedTokens := make([]string, len(tokens))

	for i, token := range tokens {
		stemmedTokens[i], _ = snowball.Stem(token, "english", false)
	}

	return stemmedTokens
}
