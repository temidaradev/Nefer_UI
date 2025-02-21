package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	log.Printf("Server is active on port %s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
