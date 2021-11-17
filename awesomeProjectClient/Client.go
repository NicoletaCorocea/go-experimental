package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/note/{id}", getNoteById)
	http.HandleFunc("/note", saveNewNote)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
