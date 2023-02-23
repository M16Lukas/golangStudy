package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
---------------------------------------------------------
* database sqlite3
---------------------------------------------------------
*/

// DB宣言
var DbConnection *sql.DB

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
                name STRING,
								age  INT)`

	// Exec コマンド実行
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

}
