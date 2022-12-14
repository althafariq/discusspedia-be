package main

import (
	"database/sql"

	"github.com/althafariq/discusspedia-be/db/migration"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "discusspedia.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	migration.Migrate(db)
}
