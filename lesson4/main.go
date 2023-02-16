package main

import "fmt"

// 定数
// 基本的にグローバルで書くことが多い

// 頭(かしら)文字を大文字にすることで、他のパッケージからも呼び出せる
// 小文字の場合は、パッケージ内でしか利用できない
const Pi = 3.14

const (
	URL      = "http://google.com"
	SiteName = "google"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 定数は最大値を設定しても、定義の段階ではエラーにならない
// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
	c0 = iota // iota : 連続する整数の連番を生成（0から）
	c1
	c2
)

func main() {
	fmt.Println(Pi) // 3.14

	// Pi = 3
	// fmt.Println(Pi)	// cannot assign to Pi

	fmt.Println(URL, SiteName) // http://google.com google

	fmt.Println(A, B, C, D, E, F) // 1 1 1 A A A

	// fmt.Println(Big - 1) // overflows
	fmt.Println(Big - 1) // 9223372036854775807

	fmt.Println(c0, c1, c2) // 0 1 2
}
