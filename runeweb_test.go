package runeweb

import (
	"reflect"
	"strings"
	"testing"

	"github.com/standupdev/runeset"
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

func TestParseLine(t *testing.T) {
	var testCases = []struct {
		line  string
		char  rune
		name  string
		words []string
	}{
		{"0021;EXCLAMATION MARK;Po;0;ON;;;;;N;;;;;",
			'!', "EXCLAMATION MARK",
			[]string{"EXCLAMATION", "MARK"}},
		{"002D;HYPHEN-MINUS;Pd;0;ES;;;;;N;;;;;",
			'-', "HYPHEN-MINUS",
			[]string{"HYPHEN", "MINUS"}},
		{"0027;APOSTROPHE;Po;0;ON;;;;;N;APOSTROPHE-QUOTE;;;",
			'\'', "APOSTROPHE (APOSTROPHE-QUOTE)",
			[]string{"APOSTROPHE", "QUOTE"}},
		{"002F;SOLIDUS;Po;0;CS;;;;;N;SLASH;;;;",
			'/', "SOLIDUS (SLASH)",
			[]string{"SOLIDUS", "SLASH"}},
	}
	for _, tc := range testCases {
		t.Run("case "+string(tc.char), func(t *testing.T) {
			rr := parseLine(tc.line)
			if rr.char != tc.char || rr.name != tc.name ||
				!reflect.DeepEqual(rr.words, tc.words) {
				t.Errorf("\nParseLine(%q)\nwant -> (%q, %q, %q)\ngot  -> (%q, %q, %q)",
					tc.line, tc.char, tc.name, tc.words, rr.char, rr.name, rr.words)
			}
		})
	}
}

const twoLines = `
003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
`

func TestBuildIndex_twoLines(t *testing.T) {
	index := BuildIndex(strings.NewReader(twoLines))
	wantChars := map[rune]string{
		'=': "EQUALS SIGN",
		'>': "GREATER-THAN SIGN",
	}
	wantWords := map[string]runeset.Set{
		"EQUALS":  runeset.Make('='),
		"GREATER": runeset.Make('>'),
		"THAN":    runeset.Make('>'),
		"SIGN":    runeset.Make('=', '>'),
	}
	if !reflect.DeepEqual(wantChars, index.Chars) {
		t.Errorf("want: %v\n got: %v", wantChars, index.Chars)
	}
	if !reflect.DeepEqual(wantWords, index.Words) {
		t.Errorf("want: %v\n got: %v", wantWords, index.Words)
	}
}

const threeLines = `
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
0043;LATIN CAPITAL LETTER C;Lu;0;L;;;;;N;;;;0063;
`

func TestBuildIndex_threeLines(t *testing.T) {
	index := BuildIndex(strings.NewReader(threeLines))
	wantChars := map[rune]string{
		'A': "LATIN CAPITAL LETTER A",
		'B': "LATIN CAPITAL LETTER B",
		'C': "LATIN CAPITAL LETTER C",
	}
	wantWords := map[string]runeset.Set{
		"A":       runeset.Make('A'),
		"B":       runeset.Make('B'),
		"C":       runeset.Make('C'),
		"LATIN":   runeset.MakeFromString("ABC"),
		"CAPITAL": runeset.MakeFromString("ABC"),
		"LETTER":  runeset.MakeFromString("ABC"),
	}
	if !reflect.DeepEqual(wantChars, index.Chars) {
		t.Errorf("want: %v\n got: %v", wantChars, index.Chars)
	}
	if !reflect.DeepEqual(wantWords, index.Words) {
		t.Errorf("want: %v\n got: %v", wantWords, index.Words)
	}
}
