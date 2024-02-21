package snippets

type TrieNode struct {
	Children map[string]*TrieNode
	End      bool
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.Root

	for _, w := range word {
		_, ok := node.Children[string(w)]
		if !ok {
			//not present, add to the child list
			newNode := &TrieNode{}
			node.Children[string(w)] = newNode
			node = node.Children[string(w)]
		} else {
			node = node.Children[string(w)]
		}
	}
	node.End = true

}

func (t *Trie) Search(word string) {

}

func (t *Trie) Prefix(prefix string) {

}

func (t *Trie) Traverse() {

}

func main() {
	t := &Trie{
		Root: &TrieNode{},
	}

	t.Insert("apple")
	t.Insert("banana")

}
