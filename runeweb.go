package runeweb

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/standupdev/runeset"
)

// Index holds a mapping of Unicode characters to names and a mapping
// of words to Unicode characters with those words in their names
type Index struct {
	Chars map[rune]string
	Words map[string]runeset.Set
}

type runeRecord struct {
	char  rune
	name  string
	words []string
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func parseLine(line string) (rr runeRecord) {
	fields := strings.Split(line, ";")
	code, err := strconv.ParseInt(fields[0], 16, 32)
	if err != nil {
		log.Fatal(err)
	}
	rr.char = rune(code)
	rr.name = fields[1]
	wordStr := strings.Replace(fields[1], "-", " ", -1)
	rr.words = strings.Fields(wordStr)
	if fields[10] != "" {
		rr.name += fmt.Sprintf(" (%s)", fields[10])
		wordStr = strings.Replace(fields[10], "-", " ", -1)
		for _, word := range strings.Fields(wordStr) {
			if !contains(rr.words, word) {
				rr.words = append(rr.words, word)
			}
		}

	}
	return rr
}

//BuildIndex builds the name index Chars and the inverted index Words
func BuildIndex(input io.Reader) (index Index) {
	index.Chars = map[rune]string{}
	index.Words = map[string]runeset.Set{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			rr := parseLine(line)
			index.Chars[rr.char] = rr.name
			for _, word := range rr.words {
				runes, found := index.Words[word]
				if found {
					runes.Add(rr.char)
				} else {
					index.Words[word] = runeset.Make(rr.char)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return index
}
