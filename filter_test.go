package runefinder

import (
	"reflect"
	"testing"

	"github.com/standupdev/runeset"
)

var index Index

func init() {
	index = BuildIndex()
}

func TestFilter(t *testing.T) {
	var testCases = []struct {
		query string
		want  runeset.Set
	}{
		{"Registered", runeset.Make('®')},
		{"ORDINAL", runeset.Make('ª', 'º')},
		{"fraction eighths", runeset.Make('⅜', '⅝', '⅞')},
		{"fraction eighths bang", runeset.Set{}},
		{"fraction eighths five", runeset.Make('⅝')},
		{"NoSuchRune", runeset.Set{}},
	}
	for _, tc := range testCases {
		t.Run(tc.query, func(t *testing.T) {
			got := Filter(index, tc.query)
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
	got := Filter(index, query)
	if len(got) < 6 || !got.Contains(want) {
		t.Errorf("query: %q\t%q absent, len(got) == %d",
			query, want, len(got))
	}
}
