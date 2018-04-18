package main

import (
	"github.com/standupdev/runefinder"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	http.HandleFunc("/", runefinder.Home)
	appengine.Main()
}
