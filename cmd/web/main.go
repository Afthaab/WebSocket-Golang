package main

import (
	"log"
	"net/http"
)

func main() {
	r := routes()
	log.Println("Starting server at 7070")
	err := http.ListenAndServe(":7070", r)
	if err != nil {
		log.Println("Error in starting the server")
	}
}
