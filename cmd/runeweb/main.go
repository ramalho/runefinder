package main

import (
	"fmt"
	"github.com/standupdev/runefinder"
	"log"
	"net/http"
)

const hostAddr = "localhost:8000"

func main() {
	fmt.Println("Serving on:", hostAddr)
	handler := http.HandlerFunc(runefinder.Home)
	log.Fatal(http.ListenAndServe(hostAddr, handler))
}
