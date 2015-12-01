package dictionary

import "unicode/utf8"

func BuildPrefixTree(d Interface) *PrefixTree {
	var root PrefixTree
	for _, w := range d.Words() {
		root.Insert(w)
	}
	return &root
}

// PrefixTree is a prefix tree/trie implementation that can be constructed from
// a dictionary.Interface.
type PrefixTree struct {
	Valid bool
	edges map[rune]*PrefixTree
}

var _ Interface = &PrefixTree{}

func (t *PrefixTree) Contains(word string) bool {
	wordBytes := []byte(word)
	for len(wordBytes) > 0 {
		if t == nil {
			return false
		}

		c, size := utf8.DecodeRune(wordBytes)
		t = t.Next(c)
		wordBytes = wordBytes[size:]
	}
	return t.Valid
}

func (t *PrefixTree) Words() (words []string) {
	if t.Valid {
		words = append(words, "")
	}

	for r, n := range t.edges {
		for _, w := range n.Words() {
			words = append(words, string(r)+w)
		}
	}
	return words
}

func (t *PrefixTree) Next(c rune) *PrefixTree {
	if t.edges == nil {
		return nil
	}
	return t.edges[c]
}

func (t *PrefixTree) Insert(s string) {
	if len(s) == 0 {
		t.Valid = true
		return
	}

	if t.edges == nil {
		t.edges = make(map[rune]*PrefixTree)
	}

	c, size := utf8.DecodeRune([]byte(s))
	rest := s[size:]

	if _, ok := t.edges[c]; !ok {
		t.edges[c] = &PrefixTree{}
	}
	t.edges[c].Insert(rest)
}
