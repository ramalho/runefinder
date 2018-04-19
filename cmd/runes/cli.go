package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/standupdev/runefinder"
	"github.com/standupdev/runeset"
)

func display(index runefinder.Index, s runeset.Set) {
	count := len(s)
	for _, c := range s.Sorted() {
		name, found := index.Chars[c]
		if !found {
			name = "(no name)"
		}
		fmt.Printf("U+%04X\t%[1]c\t%s\n", c, name)
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
	index := runefinder.BuildIndex()
	result := runefinder.Filter(index, strings.Join(os.Args[1:], " "))
	display(index, result)
}
