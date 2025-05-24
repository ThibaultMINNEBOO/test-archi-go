package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDatabase(db *sql.DB) {
	userTable := `
	CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(100),
		age INTEGER
	);
	`

	_, err := db.Exec(userTable)

	if err != nil {
		return
	}
}
