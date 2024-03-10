package textsearch

import (
	"strings"

	"github.com/kljensen/snowball"
)

func getStopWords() map[string]struct{} {
	return map[string]struct{}{
		"a": {}, "an": {}, "and": {}, "are": {}, "as": {}, "at": {},
		"be": {}, "but": {}, "by": {}, "for": {}, "if": {}, "in": {},
		"into": {}, "is": {}, "it": {}, "no": {}, "not": {}, "of": {},
		"on": {}, "or": {}, "such": {}, "that": {}, "the": {}, "their": {},
		"then": {}, "there": {}, "these": {}, "they": {}, "this": {}, "to": {},
		"was": {}, "will": {}, "with": {}, "i": {}, "have": {}, "s": {}, "my": {},
	}
}

func NormaliseFilter(tokens []string) []string {
	normalisedTokens := make([]string, len(tokens))
	for i, token := range tokens {
		// TODO: Remove all single letter words too?
		// but length of a single character maynot be neccesarily 1
		// since it might be an unicode character which is multi bytes
		// hence its len could be > 1
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
