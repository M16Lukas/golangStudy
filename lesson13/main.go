package main

import (
	"fmt"
	"golangStudy/lesson13/alib"
)

/*
------------------------------------------------------------
* テスト
* go test    		 : テスト実施
* go test ./...  : すべてのパッケージのテスト実施
* go test -v 		 : テストの詳細表示
* go test -cover : パッケージの関数のカバレッジ表示
------------------------------------------------------------
*/
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1)) // true
	fmt.Println(IsOne(0)) // false

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s)) // 3
}
