package repo

import (
	"awesomeProject/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetNoteById(id int) (model.Note, error) {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using my sql
	// The database is called notes_db
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/notes_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var note model.Note
	row := db.QueryRow("SELECT * FROM notes WHERE id = ?", id)
	defer db.Close()

	if err := row.Scan(&note.Id, &note.Name, &note.Body, &note.CreatedAt, &note.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return note, fmt.Errorf("noteById %d: no such note", id)
		}
		return note, fmt.Errorf("noteById %d: %v", id, err)
	}

	fmt.Println(note)

	return note, nil
}
