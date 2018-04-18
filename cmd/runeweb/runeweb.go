package main

import (
	"fmt"
	"github.com/standupdev/runefinder"
	"github.com/standupdev/runeset"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	hostAddr = "localhost:8000"
	form     = template.Must(template.New("form").Parse(page))
	index    = runefinder.BuildIndex()
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

func home(w http.ResponseWriter, req *http.Request) {
	chars := runeset.Set{}
	msg := ""
	query := strings.TrimSpace(req.URL.Query().Get("q"))
	if len(query) > 0 {
		chars = runefinder.Filter(index, query)
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

func main() {
	fmt.Println("Serving on:", hostAddr)
	log.Fatal(http.ListenAndServe(hostAddr, http.HandlerFunc(home)))
}

const (
	page = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Runefinder</title>
	<style>
      body {font-family: "Lucida Sans Unicode", "Lucida Grande", sans-serif;}
	  table {font-family: "Lucida Console", Monaco, monospace; text-align: left; min-width: 300px}
      td.code {min-width: 40px; text-align: right;}
      td.char {min-width: 50px; text-align: center;}
	  caption {background: lightgray; }
	</style>
  </head>
  <body>
    <p>
      <form action="/">
        <input type="search" name="q" value="{{.Query}}">
        <input type="submit" value="find">
	    Examples:
		{{range .Links}}
			<a href="{{.Location}}" 
               title="find &quot;{{.Text}}&quot;">{{.Text}}</a>
		{{end}}
      </form>
    </p>

    <table>
      <caption>{{.Message}}</caption>
      {{range .Result}}
        <tr>
          <td class="code">{{.Code}}</td>
          <td class="char">{{.Char}}</td>
          <td>{{.Name}}</td>
        </tr>
      {{end}}
    </table>
  </body>
</html>
`
	sampleWords = `
bismillah
box
cat
chess
circle
circled
Egyptian
face
hexagram
key
Malayalam
operator
Roman
symbol`
)
