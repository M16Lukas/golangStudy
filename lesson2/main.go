package main

import "fmt"

// 変数
// 注意：定義された変数は必ずプログラム上で使う必要がある

func main() {
	// 明示的な定義（基本的にこれを使うことをお勧め）
	// 関数外でも使える
	// var 変数名　型 =　値
	var i int = 200
	fmt.Println(i) // 200

	var s string = "hello go"
	fmt.Println(s) // hello go

	var t, f bool = true, false
	fmt.Println(t, f) // true false

	var (
		i2 int    = 300
		s2 string = "Bye go"
	)
	fmt.Println(i2, s2) // 300 Bye go

	// 変数だけを指定
	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 0 空文字

	i3 = 340
	s3 = "golang"
	fmt.Println(i3, s3) // 340 golang

	// 値の更新
	i = 150
	fmt.Println(i) // 150

	// 暗黙敵な定義
	// 注意：関数内でしか定義できない
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4) // 400
}
