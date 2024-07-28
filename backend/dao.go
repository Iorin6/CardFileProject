package backend

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type BookJSON struct {
	ID          int64
	name        string
	author      string
	year        string
	genre       string
	description string
}

func GetAllBooks(db *sql.DB) []BookJSON {
	rows, err := db.Query(`SELECT "id", "name", "author", "year", "genre", "description" FROM "books"`)
	CheckError(err)
	return rowToBookJSON(rows)
}

func GetBookByName(name string, db *sql.DB) []BookJSON {
	rows, err := db.Query(`SELECT "id", "name", "author", "year", "genre", "description" FROM "books" where name = $1`, name)
	CheckError(err)
	return rowToBookJSON(rows)
}
