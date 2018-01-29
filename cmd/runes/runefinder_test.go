package main

import (
	"os"
	"reflect"
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
		query string
		want  runeset.Set
	}{
		{"Registered", runeset.Make('®')},
		{"ORDINAL", runeset.Make('ª', 'º')},
		{"fraction eighths", runeset.Make('⅜', '⅝', '⅞')},
		{"fraction eighths five", runeset.Make('⅝')},
		{"NoSuchRune", runeset.Set{}},
	}
	for _, tc := range testCases {
		t.Run(tc.query, func(t *testing.T) {
			got := filter(index, tc.query)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("query: %q\twant: %q\tgot: %q",
					tc.query, tc.want, got)
			}
		})
	}
}

func TestFilter_hyphenatedQuery(t *testing.T) {
	query := "HYPHEN-MINUS"
	want := '-'
	got := filter(index, query)
	if len(got) < 6 || !got.Has(want) {
		t.Errorf("query: %q\t%q absent, len(got) == %d",
			query, want, len(got))
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

func Example_single_result() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "registered"}
	main()
	// Output:
	// U+00AE	®	REGISTERED SIGN (REGISTERED TRADE MARK SIGN)
	// 1 character found
}

func Example_no_result() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "aintnocharacterlikethis"}
	main()
	// Output:
	// no character found
}
