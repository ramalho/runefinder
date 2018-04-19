package runefinder

import (
	"fmt"
	"github.com/standupdev/runeset"
	"html/template"
	"net/http"
	"strings"
)

const (
	sampleWords = `bismillah box cat chess circle circled 
                   Egyptian face hexagram key Malayalam Roman symbol`
	formPath = "data/form.html"
)

var (
	formHTML = string(MustAsset(formPath))
	form     = template.Must(template.New("form").Parse(formHTML))
	index    = BuildIndex()
	links    = makeLinks(sampleWords)
)

type Link struct {
	Location template.URL
	Text     string
}

func makeLinks(text string) []Link {
	links := []Link{}
	for _, word := range strings.Fields(text) {
		location := template.URL("/?q=" + word)
		links = append(links, Link{location, word})
	}
	return links
}

func makeMessage(count int) string {
	switch count {
	case 0:
		return "No character found."
	case 1:
		return "1 character found."
	default:
		return fmt.Sprintf("%d characters found", count)
	}
}

func getName(char rune) string {
	name, found := index.Chars[char]
	if !found {
		name = "(no name)"
	}
	return name
}

type RuneRecord struct {
	Code string
	Char string
	Name string
}

func makeResults(chars runeset.Set) []RuneRecord {
	result := []RuneRecord{}
	for _, char := range chars.Sorted() {
		result = append(result, RuneRecord{
			Code: fmt.Sprintf("U+%04X", char),
			Char: string(char),
			Name: getName(char),
		})
	}
	return result
}

func Home(w http.ResponseWriter, req *http.Request) {
	chars := runeset.Set{}
	msg := ""
	query := strings.TrimSpace(req.URL.Query().Get("q"))
	if len(query) > 0 {
		chars = Filter(index, query)
		msg = makeMessage(len(chars))
	}
	data := struct {
		Links   []Link
		Query   string
		Message string
		Result  []RuneRecord
	}{
		Links:   links,
		Query:   query,
		Message: msg,
		Result:  makeResults(chars),
	}
	form.Execute(w, data)
}
