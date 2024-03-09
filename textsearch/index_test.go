package textsearch_test

import (
	"practice/textsearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexAdd(t *testing.T) {
	index := make(textsearch.Index)

	assert.Nil(t, index.Search("hello there"))
	assert.Nil(t, index.Search("shazam!!"))

}
