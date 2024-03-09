package textsearch

import (
	"fmt"
	"os"
)

func ExampleSearch() {
	invertedIndex := make(Index)
	sampleDoc1 := Document{
		ID:    1,
		Title: "Anarchism of my cat",
		Text:  "This is my cat. His name is Anakin. He's a one little naught boy.",
		URL:   "https://wiki.cat.example",
	}

	sampleDoc2 := Document{
		ID:    2,
		Title: "Obedience of my dog",
		Text:  "This is my dog. His name is Anakin. He's one obedient boy.",
		URL:   "https://wiki.dog.example",
	}

	sampleDoc3 := Document{
		ID:    3,
		Title: "Anarchism of wikipedia",
		Text: `Anarchism is a political philosophy and movement that is skeptical of all justifications for authority
		and seeks to abolish the institutions it claims maintain unnecessary coercion and hierarchy, typically including nation-states, 
		and capitalism. Anarchism advocates for the replacement of the state with stateless societies and voluntary free associations.`,
		URL: "https://wiki.idx.example",
	}

	docs := make([]Document, 3)
	docs[0] = sampleDoc1
	docs[1] = sampleDoc2
	docs[2] = sampleDoc3

	invertedIndex.Add(docs)

	query := "anakin is a good dog!"
	matchedIDs := invertedIndex.Search(query)
	if len(matchedIDs) == 0 {
		fmt.Printf("no docs found matching the query: %s", query)
		os.Exit(1)
	}

	for _, id := range matchedIDs {
		doc := docs[id-1]
		fmt.Printf("doc ID: %d and doc Text: %s\n", doc.ID, doc.Text)
	}

}
