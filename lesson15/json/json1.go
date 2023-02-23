package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/*
------------------------------------------
* JSON
------------------------------------------
*/

// 構造体からJSONテキストへの変換

type A struct{}

type User struct {
	ID      int       `json:"id,string"` // jsonに変換する時、型をstringに変更
	Name    string    `json:"-"`         // - :JSONに変換する時、表示しない
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A",omitempty` // omitempty : フィールドの初期値の場合、隠す
}

// UnmarshalJSON : Unmarshalのカスタム
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

// MarshalJSON : Marshalのカスタム
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	// Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)         // [123 34 78 97 109 101 34 58 34 77 114 32 116 101 115 116 34 125]
	fmt.Println(string(bs)) // {"Name":"Mr test"}

	// ------------------------------

	fmt.Printf("%T\n", bs) // []uint8

	u2 := new(User)

	fmt.Printf("u2の型： %T\n", u2) // u2の型： *main.User

	// Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2) // {0 Mr test!  0001-01-01 00:00:00 +0000 UTC <nil>}
}
