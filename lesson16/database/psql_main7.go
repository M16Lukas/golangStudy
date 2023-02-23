package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ = sql.Open("postgres", "user=test_user dbname=test_db password=password sslmode=disable")
	defer Db.Close()

	// C

	// R

	// U

	// D

}
