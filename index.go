package runefinder

import (
	"strings"

	"github.com/standupdev/runeset"
	"golang.org/x/text/unicode/runenames"
)

// Index is an inverted index: a mapping of words to
// Unicode characters with those words in their names
type Index map[string]runeset.Set

const (
	firstChar rune = 0x20
	lastChar  rune = 0x10FFFF // http://unicode.org/faq/utf_bom.html
)

func parseName(name string) []string {
	name = strings.Replace(name, "-", " ", -1)
	words := []string{}
	for _, word := range strings.Fields(name) {
		words = append(words, word)
	}
	return words
}

//buildIndex builds the inverted index for the given range of runes
func buildIndex(first, last rune) (index Index) {
	index = Index{}
	for char := first; char <= last; char++ {
		name := runenames.Name(char)
		if len(name) > 0 {
			for _, word := range parseName(name) {
				runes, found := index[word]
				if found {
					runes.Add(char)
				} else {
					index[word] = runeset.Make(char)
				}
			}
		}
	}
	return index
}

//BuildIndex builds the inverted index for all runes
func BuildIndex() (index Index) {
	return buildIndex(firstChar, lastChar)
}
