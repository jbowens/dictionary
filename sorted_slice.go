package dictionary

import "sort"

type sortedSlice []string

// Assert that sortedSlice implements dictionary.Interface.
var _ Interface = sortedSlice(nil)

// WithWords creates a new dictionary with the provided words.
func WithWords(words ...string) Interface {
	dict := sortedSlice(words)
	dict.sortIfNecessary()
	return dict
}

// Contains determines if the provided word is contained within the dictionary.
func (s sortedSlice) Contains(w string) bool {
	i := sort.SearchStrings(s, w)
	if i >= len(s) {
		return false
	}

	return s[i] == w
}

// Words implements dictionary.Interface.
func (s sortedSlice) Words() []string {
	return []string(s)
}

func (s sortedSlice) sortIfNecessary() {
	// Sort the words if they're not already sorted. Most of the time the words should
	// already be sorted if they're coming from /usr/share/dict/words.
	if !sort.StringsAreSorted(s) {
		sort.Strings(s)
	}
}
