package textsearch_test

import (
	"practice/textsearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalisedTokens(t *testing.T) {
	testCases := []struct {
		input  []string
		result []string
	}{
		{
			input:  []string{"WgGres", "my", "DoG"},
			result: []string{"wggres", "my", "dog"},
		},
	}

	for _, testCase := range testCases {
		t.Run("normalisation tests", func(t *testing.T) {
			assert.EqualValues(t, testCase.result, textsearch.NormaliseFilter(testCase.input))
		})
	}
}

func TestStopWords(t *testing.T) {
	testCases := []struct {
		input  []string
		result []string
	}{
		{
			input:  []string{"i", "am", "the", "cat"},
			result: []string{"am", "cat"},
		},
	}

	for _, testCase := range testCases {
		t.Run("stopwords tests", func(t *testing.T) {
			assert.EqualValues(t, testCase.result, textsearch.StopWordsFilter(testCase.input))
		})
	}
}

func TestStemming(t *testing.T) {
	testCases := []struct {
		input  []string
		result []string
	}{
		{
			input:  []string{"cat", "cats", "fish", "fishing", "fished", "airline"},
			result: []string{"cat", "cat", "fish", "fish", "fish", "airlin"},
		},
	}

	for _, testCase := range testCases {
		t.Run("stemming tests", func(t *testing.T) {
			assert.EqualValues(t, testCase.result, textsearch.StemmingFilter(testCase.input))
		})
	}
}
