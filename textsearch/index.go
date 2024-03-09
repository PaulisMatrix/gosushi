package textsearch

// inverted index to map each word to all the document IDs it occurs in
type Index map[string][]int

func (idx Index) Add(docs []Document) {

}

func (idx Index) Search(query string) []int {
	return nil
}
