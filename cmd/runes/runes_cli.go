package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/standupdev/runeset"
	"github.com/standupdev/runeweb"
)

var logger = log.New(os.Stderr, "", log.Lshortfile)

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
	index := runeweb.ReadIndex()
	result := runeweb.Filter(index, strings.Join(os.Args[1:], " "))
	display(index, result)
}
