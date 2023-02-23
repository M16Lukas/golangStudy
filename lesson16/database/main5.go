package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// データ追加
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := DbConnection.Exec(cmd, "hanako", 19)
	if err != nil {
		log.Fatalln(err)
	}

	// 特定のデータを取得
	cmd2 := "SELECT * FROM persons where age = ?"
	// QueryRow：１レコード取得
	row := DbConnection.QueryRow(cmd2, 19)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age) // hanako 19
}
