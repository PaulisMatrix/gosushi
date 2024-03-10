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

	index.Add([]textsearch.Document{{ID: 1, Text: "A cat ran over my keyboard"}})
	assert.EqualValues(t, 1, index.WordFreq("cat"))
	assert.EqualValues(t, []int{1}, index.Search("my cat"))

	index.Add([]textsearch.Document{{ID: 2, Text: "my cat's name is anakin"}})
	assert.EqualValues(t, []int{}, index.Search("a"))
	assert.EqualValues(t, 2, index.WordFreq("cat"))
	assert.EqualValues(t, []int{1, 2}, index.Search("Cat"))
	assert.EqualValues(t, 1, index.WordFreq("anakin"))

	index.Add([]textsearch.Document{{ID: 3, Text: "my cat is a real menace"}})
	assert.EqualValues(t, 3, index.WordFreq("cat"))
	assert.EqualValues(t, []int{1, 2, 3}, index.Search("how is my cat?"))

}

func TestIntersection(t *testing.T) {
	set1 := []int{100, 90, 24, 1, 0, 12, 44, 646, 34366, 412}
	set2 := []int{0, 1, 100, 24, 93130, 8731, 34, 12}
	set3 := []int{100, 90, 78, 12, 900, 78912, 12122}

	intersection := make([]int, 0)

	for i := 0; i < 2; i++ {
		if len(intersection) == 0 {
			intersection = textsearch.Intersection(set1, set2)
		} else {
			intersection = textsearch.Intersection(intersection, set3)
		}
	}

	assert.EqualValues(t, []int{100, 12}, intersection)

}
