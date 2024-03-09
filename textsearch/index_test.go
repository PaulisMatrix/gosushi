package textsearch_test

import (
	"practice/textsearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexAdd(t *testing.T) {
	index := make(textsearch.Index)

	assert.Equal(t, []int{}, index.Search("hello there"))
	assert.Equal(t, []int{}, index.Search("shazam!!"))

	index.Add([]textsearch.Document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	assert.Equal(t, []int{}, index.Search("a"))
	assert.Equal(t, index.Search("donut"), []int{1})
	assert.Equal(t, index.Search("DoNuts"), []int{1})
	assert.Equal(t, index.Search("glass"), []int{1})

	index.Add([]textsearch.Document{{ID: 2, Text: "donut is a donut"}})
	assert.Equal(t, []int{}, index.Search("a"))
	assert.Equal(t, index.Search("donut"), []int{1, 2})
	assert.Equal(t, index.Search("DoNuts"), []int{1, 2})
	assert.Equal(t, index.Search("glass"), []int{1})

}
