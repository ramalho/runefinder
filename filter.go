package runeweb

import (
	"strings"

	"github.com/standupdev/runeset"
)

func Filter(index Index, query string) (result runeset.Set) {
	query = strings.Replace(query, "-", " ", -1)
	query = strings.ToUpper(query)
	words := strings.Fields(query)
	result, found := index.Words[words[0]]
	if !found {
		return runeset.Set{}
	}
	for _, word := range words[1:] {
		chars, found := index.Words[word]
		if !found {
			return runeset.Set{}
		}
		result = result.Intersection(chars)
		if len(result) == 0 {
			break
		}
	}
	return result
}
