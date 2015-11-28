package dictionary

type inMemory struct {
	words map[string]struct{}
}

// Assert that inMemory implements dictionary.Interface.
var _ Interface = &inMemory{}

// WithWords creates a new dictionary with the provided words.
func WithWords(words ...string) Interface {
	dict := &inMemory{
		words: make(map[string]struct{}),
	}

	for _, w := range words {
		dict.words[w] = struct{}{}
	}

	return dict
}

// Contains determines if the provided word is contained within the dictionary.
func (d *inMemory) Contains(w string) bool {
	_, ok := d.words[w]
	return ok
}

// Words implements dictionary.Interface.
func (d *inMemory) Words() []string {
	words := make([]string, 0, len(d.words))

	for k := range d.words {
		words = append(words, k)
	}

	return words
}
