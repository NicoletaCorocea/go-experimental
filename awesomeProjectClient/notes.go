package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type NotesResponse struct {
	Note []Note `json:"notes"`
}

type Note struct {
	id        int       `json:"id"`
	name      string    `json:"name"`
	body      string    `json:"body"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
}

func getNoteById(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8081/note/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "I'm the client"+string(responseData), html.EscapeString(r.URL.Path))
}

func saveNewNote(w http.ResponseWriter, r *http.Request) {

}

func deleteNoteById(w http.ResponseWriter, r *http.Request) {

}
