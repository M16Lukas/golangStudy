package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	//C
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	_, err := Db.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}

	//R
	cmd2 := "SELECT * FROM persons where age = $1"
	// QueryRow １レコード取得
	row := Db.QueryRow(cmd2, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		// データがなかったら
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age) // Nancy 20

	cmd = "SELECT * FROM persons"
	// Query は条件に合うものをすべて取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	// structを作成
	var pp []Person
	// rows.Next()　取得したデータをループでスライスに追加
	for rows.Next() {
		var p Person
		// scan データ追加
		err := rows.Scan(&p.Name, &p.Age)
		// １つずつエラーハンドリング
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age) // Nancy 20
	}

	//U
	cmd3 := "UPDATE persons SET age = $1 WHERE name = $2"
	_, err2 := Db.Exec(cmd3, 25, "Nancy")
	if err2 != nil {
		log.Fatalln(err2)
	}

	//D
	cmd4 := "DELETE FROM persons WHERE name = $1"
	_, err3 := Db.Exec(cmd4, "Nancy")
	if err3 != nil {
		log.Fatalln(err3)
	}

}
