package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/standupdev/runeweb"
	"strings"
	"github.com/standupdev/runeset"
	"bufio"
	"strconv"
)

const (
	unicodeDataURL  = "http://www.unicode.org/Public/UNIDATA/UnicodeData.txt"
	unicodeDataPath = "UnicodeData.txt"
	indexPath       = "runeweb_index.gob"
)

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

//buildIndex builds the name index Chars and the inverted index Words
func buildIndex(input io.Reader) (index runeweb.Index) {
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


func saveIndex(index runeweb.Index, indexPath string) {
	indexFile, err := os.Create(indexPath)
	if err != nil {
		log.Fatal("Unable to create index file.")
	} else {
		defer indexFile.Close()
		encoder := gob.NewEncoder(indexFile)
		err := encoder.Encode(index)
		if err != nil {
			log.Fatal("encode error:", err)
		}
	}
}

func failIf(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func fetchUCD(url, path string, done chan<- bool) {
	response, err := http.Get(url)
	failIf(err)
	defer response.Body.Close()
	file, err := os.Create(path)
	failIf(err)
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	failIf(err)
	done <- true
}

func progress(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println()
			return
		default:
			fmt.Print(".")
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func openUnicodeData(path string) (*os.File, error) {
	ucd, err := os.Open(path)
	if os.IsNotExist(err) {
		fmt.Printf("%s not found\ndownloading %s\n", path, unicodeDataURL)
		done := make(chan bool)
		go fetchUCD(unicodeDataURL, path, done)
		progress(done)
		ucd, err = os.Open(path)
	}
	return ucd, err
}

func main() {
	ucd, err := openUnicodeData(unicodeDataPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ucd.Close()
	saveIndex(buildIndex(ucd), indexPath)
}
