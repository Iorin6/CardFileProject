package backend

import (
	"database/sql"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func rowToBookJSON(rows *sql.Rows) []BookJSON {
	var books []BookJSON
	for rows.Next() {
		var id int64
		var name string
		var author string
		var year string
		var genre string
		var description string

		err := rows.Scan(&id, &name, &author, &year, &genre, &description)
		CheckError(err)

		books = append(books, BookJSON{ID: id, name: name, author: author, year: year, genre: genre, description: description})
	}
	return books
}
