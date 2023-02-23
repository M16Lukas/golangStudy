package main

import (
	"database/sql"
	"fmt"
	"log"

	//インポート _にしないとコンパイルエラーになる。使用しない為。ビルドして一緒にコンパイルしないとsqlにアクセスできない為。
	_ "github.com/mattn/go-sqlite3"
)

//データベース操作 sqlite3 ver
//テーブル作成

// コネクションプール　グローバルに宣言
var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// 複数のデータの取得
	cmd := "SELECT * FROM persons"
	// Query は条件に合うものをすべて取得
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
	// Nancy 20
	// hanako 19
}
