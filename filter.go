package runeweb

import (
	"strings"

	"github.com/standupdev/runeset"
)

func Filter(index Index, query string) (result runeset.Set) {
	query = strings.Replace(query, "-", " ", -1)
	words := strings.Fields(query)
	for i, word := range words {
		word = strings.ToUpper(word)
		chars, found := index.Words[word]
		if !found {
			return runeset.Set{}
		}
		if i == 0 {
			result = chars
		} else {
			result = result.Intersection(chars)
			if len(result) == 0 {
				break
			}
		}
	}
	return result
}
