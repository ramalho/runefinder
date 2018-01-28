package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/standupdev/runeset"
	"github.com/standupdev/runeweb"
)

const indexPath = "runeweb_index.gob"

var logger = log.New(os.Stderr, "", log.Lshortfile)

func readIndex(path string) (index runeweb.Index) {
	indexFile, err := os.Open(path)
	if err != nil {
		logger.Fatalln(err)
	}
	defer indexFile.Close()
	decoder := gob.NewDecoder(indexFile)
	err = decoder.Decode(&index)
	if err != nil {
		logger.Fatalln(err)
	}
	return index
}

func filter(index runeweb.Index, query string) (result runeset.Set) {
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

func display(index runeweb.Index, s runeset.Set) {
	count := 0
	for _, c := range s.Sorted() {
		name, found := index.Chars[c]
		if !found {
			name = "(no name)"
		}
		fmt.Printf("U+%04X\t%[1]c\t%s\n", c, name)
		count++
	}
	fmt.Println(count, "characters found")
}

func main() {
	index := readIndex(indexPath)
	result := filter(index, strings.Join(os.Args[1:], " "))
	display(index, result)
}
