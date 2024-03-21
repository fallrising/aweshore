package model

import "time"

// Note represents a note record in the database
type Note struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
