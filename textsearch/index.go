package textsearch

const MAX_BIT_LEN int = 31
const MAX_BITMAP_LEN int = 22581

// inverted index to map each word to all the document IDs it occurs in
type Index map[string][]int

func (idx Index) Add(docs []Document) {

	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			// what if cat is present at two places in the same text
			// how would you avoid adding the same doc ID twice?
			curIds := idx[token]
			if curIds != nil && curIds[len(curIds)-1] == doc.ID {
				continue
			}
			idx[token] = append(curIds, doc.ID)
		}
	}

}

func (idx Index) Search(query string) []int {
	docIDs := make([]int, 0)

	for _, token := range analyze(query) {
		// get the docIDs list from inverted index for each token
		// find the common IDs from all such list
		ids, ok := idx[token]
		if !ok {
			// token doesn't exist, do we just return or return the found docIDs
			continue
		}
		if len(docIDs) == 0 {
			// init
			docIDs = ids
			continue
		}
		docIDs = intersection(docIDs, ids)
	}

	return docIDs
}

// using bitmaps to find out the common ids between the two sets
// length of the docs(in this case) ~7_00_000
// max integer I can store in a bitmap of length 1 is 31 (index 0 to 31 for a 32 bit integer)
// so here, I would be needing 7k / 31 len bitmap ~ 226

func intersection(seta, setb []int) []int {
	intersection := make([]int, 0)
	// arbitarily taken length
	bitmap := make([]int, MAX_BITMAP_LEN)

	// compute the bitmap first
	for _, item := range seta {
		// index to know where should we put the integer
		// mask to calculate the actual mask
		// example : https://go.dev/play/p/Wv-9Wxn-4Lj
		index, mask := divmod(item, MAX_BIT_LEN)
		bitMask := int(1 << mask)
		bitmap[index] |= bitMask
	}

	// calculate the intersection
	for _, item := range setb {
		index, mask := divmod(item, MAX_BIT_LEN)
		source := int(1 << mask)
		target := bitmap[index]

		if source&target != 0 {
			intersection = append(intersection, item)
		}
	}

	return intersection

}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}
