package textsearch

const MAX_BIT_LEN uint32 = 31
const MAX_BITMAP_LEN uint32 = 226

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
	return nil
}

// using bitmaps to find out the common ids between the two sets
// length of the docs(in this case) ~7000
// max integer I can store in a bitmap of length 1 is 31 (index 0 to 31 for a 32 bit integer)
// so here, I would be needing 7k / 31 len bitmap ~ 226

func intersection(seta, setb []uint32) []uint32 {
	intersection := make([]uint32, 0)
	// arbitarily taken length
	bitmap := make([]uint32, MAX_BITMAP_LEN)

	// compute the bitmap first
	for _, item := range seta {
		// index to know where should we put the integer
		// mask to calculate the actual mask
		// example : https://go.dev/play/p/Wv-9Wxn-4Lj
		index, mask := divmod(item, MAX_BIT_LEN)
		bitMask := uint32(1 << mask)
		bitmap[index] |= bitMask
	}

	// calculate the intersection
	for _, item := range setb {
		index, mask := divmod(item, MAX_BIT_LEN)
		source := uint32(1 << mask)
		target := bitmap[index]

		if source&target != 0 {
			intersection = append(intersection, item)
		}
	}

	return intersection

}

func divmod(numerator, denominator uint32) (quotient, remainder uint32) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}
