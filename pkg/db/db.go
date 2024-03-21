// File: pkg/db/db.go
package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "aweshore.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable()
}

func createTable() {
	createNoteTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"content" TEXT,
		"created" DATETIME,
		"updated" DATETIME
	  );`

	statement, err := DB.Prepare(createNoteTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}
