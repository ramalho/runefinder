package runeweb

import (
	"log"
	"strings"

	"github.com/standupdev/runeset"
	"encoding/gob"
	"bytes"
	"os"
)

var logger = log.New(os.Stderr, "", log.Lshortfile)

// Index holds a mapping of Unicode characters to names and a mapping
// of words to Unicode characters with those words in their names
type Index struct {
	Chars map[rune]string
	Words map[string]runeset.Set
}


const indexPath = "data/runeweb_index.gob"

func ReadIndex() (index Index) {
	indexData, err := Asset(indexPath)
	if err != nil {
		logger.Fatalln(err)
	}
	decoder := gob.NewDecoder(bytes.NewReader(indexData))
	err = decoder.Decode(&index)
	if err != nil {
		logger.Fatalln(err)
	}
	return index
}

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
