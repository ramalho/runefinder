package runefinder

import (
	"strings"

	"github.com/standupdev/runeset"
)

func Filter(index Index, query string) (result runeset.Set) {
	query = strings.Replace(query, "-", " ", -1)
	query = strings.ToUpper(query)
	words := strings.Fields(query)
	chars, found := index[words[0]]
	if !found {
		return runeset.Set{}
	}
	result = chars.Copy()
	for _, word := range words[1:] {
		chars, found := index[word]
		if !found {
			return runeset.Set{}
		}
		result.IntersectionUpdate(chars)
		if len(result) == 0 {
			break
		}
	}
	return result
}
