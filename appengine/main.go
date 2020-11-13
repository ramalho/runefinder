package main

import (
	"log"
	"net/http"
	"os"

	"github.com/standupdev/runefinder"
)

func main() {
	http.HandleFunc("/", runefinder.Home)
	port := getPort()
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
