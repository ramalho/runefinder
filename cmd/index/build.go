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
)

const (
	unicodeDataURL  = "http://www.unicode.org/Public/UNIDATA/UnicodeData.txt"
	unicodeDataPath = "UnicodeData.txt"
	indexPath       = "runeweb_index.gob"
)

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
	saveIndex(runeweb.BuildIndex(ucd), indexPath)
}
