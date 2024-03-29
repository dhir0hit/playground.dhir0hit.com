package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"playground.dhir0hit.com/Controller"
)

func main() {
	// App Entry point
	Controller.AppEntry()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
