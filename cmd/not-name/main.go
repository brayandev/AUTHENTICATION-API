package main

import (
	"log"
	"net/http"
)

func main() {
	handler, err := createServerHandler()
	if err != nil {
		log.Fatal("Failed to start server")
	}
	http.ListenAndServe(":8080", handler)
}
