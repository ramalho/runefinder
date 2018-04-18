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
	log.Fatal(http.ListenAndServe(hostAddr, http.HandlerFunc(runefinder.Home)))
}
