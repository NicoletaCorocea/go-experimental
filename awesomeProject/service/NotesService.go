package service

import (
	"awesomeProject/model"
	"awesomeProject/repo"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])

	var note model.Note
	note, _ = repo.GetNoteById(id)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	order := r.FormValue("order")

	fmt.Println(order)

	var note1 model.Note
	note1.Id = 1

	var note2 model.Note
	note2.Id = 2

	var notes []model.Note
	notes = append(notes, note1)
	notes = append(notes, note2)

	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(notes)
	if err != nil {
		return
	}

	//update response
	w.Write(jsonResponse)

	//json.NewEncoder(w).Encode(note)

}

func SaveNewNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note model.Note
	_ = json.NewDecoder(r.Body).Decode(note)
	note.Id = rand.Intn(1000000)

	fmt.Println(note)
	json.NewEncoder(w).Encode(&note)
}

