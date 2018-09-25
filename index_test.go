package runefinder

import (
	"github.com/standupdev/runeset"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	var testCases = []struct {
		haystack []string
		needle   string
		want     bool
	}{
		{[]string{"EXCLAMATION", "MARK"}, "MARK", true},
		{[]string{"EXCLAMATION", "MARK"}, "BEAR", false},
		{[]string{}, "", false},
	}
	for _, tc := range testCases {
		t.Run(tc.needle, func(t *testing.T) {
			if got := contains(tc.haystack, tc.needle); tc.want != got {
				t.Errorf("contains(%q, %q)\twant -> %t\tgot  -> %t",
					tc.haystack, tc.needle, tc.want, got)
			}
		})
	}
}

func TestParseName(t *testing.T) {
	var testCases = []struct {
		name  string
		words []string
	}{
		{"EXCLAMATION MARK",
			[]string{"EXCLAMATION", "MARK"}},
		{"HYPHEN-MINUS",
			[]string{"HYPHEN", "MINUS"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseName(tc.name)
			if !reflect.DeepEqual(got, tc.words) {
				t.Errorf("\nParseName(%q)\nwant -> %q\ngot  -> %q",
					tc.name, tc.words, got)
			}
		})
	}
}

func TestBuildIndex_twoLines(t *testing.T) {
	// 003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
	// 003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
	index := buildIndex(0x3D, 0x3E)
	wantWords := Index{
		"EQUALS":  runeset.Make('='),
		"GREATER": runeset.Make('>'),
		"THAN":    runeset.Make('>'),
		"SIGN":    runeset.Make('=', '>'),
	}
	if !reflect.DeepEqual(wantWords, index) {
		t.Errorf("want: %v\n got: %v", wantWords, index)
	}
}

func TestBuildIndex_threeLines(t *testing.T) {
	// 0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
	// 0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
	// 0043;LATIN CAPITAL LETTER C;Lu;0;L;;;;;N;;;;0063;
	index := buildIndex(0x41, 0x43)
	wantWords := Index{
		"A":       runeset.Make('A'),
		"B":       runeset.Make('B'),
		"C":       runeset.Make('C'),
		"LATIN":   runeset.MakeFromString("ABC"),
		"CAPITAL": runeset.MakeFromString("ABC"),
		"LETTER":  runeset.MakeFromString("ABC"),
	}
	if !reflect.DeepEqual(wantWords, index) {
		t.Errorf("want: %v\n got: %v", wantWords, index)
	}
}

var registeredSign rune = 0xAE // ®

func TestUnicodeDataIndex_Words(t *testing.T) {
	index := BuildIndex()
	wantWords := 10000
	if len(index) < wantWords {
		t.Errorf("len(index.Chars) < %d\t got: %d", wantWords, len(index))
	}
	wantSet := runeset.Make(registeredSign)
	gotSet := index["REGISTERED"]
	if !reflect.DeepEqual(wantSet, gotSet) {
		t.Errorf("want: %v\t got: %v", wantSet, gotSet)
	}
}
