package main

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/standupdev/runeset"
	"github.com/standupdev/runeweb"
)

var index runeweb.Index

func init() {
	index = readIndex(indexPath)
}

func TestFilter(t *testing.T) {
	var testCases = []struct {
		query []string
		want  runeset.Set
	}{
		{[]string{"ordinal"}, runeset.Make('ª', 'º')},
		{[]string{"fraction", "eighths"}, runeset.Make('⅜', '⅝', '⅞')},
		{[]string{"NoSuchRune"}, runeset.Set{}},
	}
	for _, tc := range testCases {
		label := strings.Join(tc.query, " ")
		t.Run(label, func(t *testing.T) {
			got := filter(index, tc.query)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("query: %q\twant: %q\tgot: %q",
					tc.query, tc.want, got)
			}
		})
	}
}

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "EIGHTHS", "fraction"}
	main()
	// Output:
	// U+215C	⅜	VULGAR FRACTION THREE EIGHTHS (FRACTION THREE EIGHTHS)
	// U+215D	⅝	VULGAR FRACTION FIVE EIGHTHS (FRACTION FIVE EIGHTHS)
	// U+215E	⅞	VULGAR FRACTION SEVEN EIGHTHS (FRACTION SEVEN EIGHTHS)
	// 3 characters found
}
