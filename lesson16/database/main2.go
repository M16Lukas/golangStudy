package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// データ追加
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := DbConnection.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}
}
