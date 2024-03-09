package textsearch_test

import (
	"practice/textsearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		text   string
		tokens []string
	}{
		{
			text:   "I have a small;cat",
			tokens: []string{"I", "have", "a", "small", "cat"},
		},
		{
			text:   "",
			tokens: []string{},
		},
		{
			text:   "\r\tthis\bis;wild\t",
			tokens: []string{"this", "is", "wild"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.text, func(t *testing.T) {
			assert.EqualValues(t, testCase.tokens, textsearch.Tokenize(testCase.text))
		})
	}
}
