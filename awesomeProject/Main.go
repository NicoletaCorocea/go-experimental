package main

import (
	"awesomeProject/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/note/{id}", service.GetNoteById).Methods("GET")
	router.HandleFunc("/note", service.GetAllNotes).Queries("order", "{order}").Methods("GET")
	router.HandleFunc("/note", service.SaveNewNote).Methods("POST")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
