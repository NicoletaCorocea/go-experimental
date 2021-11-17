package model

import (
	"time"
)

type NotesResponse struct {
	Note []Note `json:"notes"`
}

// the use of capital letter makes it exported automatically. It is similar to public access
// lower case is like a private
type Note struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
