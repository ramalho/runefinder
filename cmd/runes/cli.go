package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/standupdev/runefinder"
	"github.com/standupdev/runeset"
	"golang.org/x/text/unicode/runenames"
)

func display(s runeset.Set) {
	count := len(s)
	for _, c := range s.Sorted() {
		name := runenames.Name(c)
		fmt.Printf("%U\t%[1]c\t%s\n", c, name)
	}
	var msg string
	switch count {
	case 0:
		msg = "no character found"
	case 1:
		msg = "1 character found"
	default:
		msg = fmt.Sprintf("%d characters found", count)
	}
	fmt.Println(msg)
}

func main() {
	query := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	if len(query) == 0 {
		fmt.Println("Please provide one or more words or characters to search.")
		return
	}
	nameChars := runeset.MakeFromString(" -0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	charQuery := false
	for _, char := range query {
		if !nameChars.Contains(char) {
			charQuery = true
			break
		}
	}
	if charQuery {
		display(runeset.MakeFromString(query))
	} else {
		index := runefinder.BuildIndex()
		display(runefinder.Filter(index, query))
	}
}
