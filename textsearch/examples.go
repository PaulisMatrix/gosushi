package textsearch

import "fmt"

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

	invertedIndex.Add([]Document{sampleDoc1})
	invertedIndex.Add([]Document{sampleDoc2})
	invertedIndex.Add([]Document{sampleDoc3})
	fmt.Printf("sample inverted index: %+v\n", invertedIndex)

	// now if I search for anakin, it should return the docs 1 and 2
}
