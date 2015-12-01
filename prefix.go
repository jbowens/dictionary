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
