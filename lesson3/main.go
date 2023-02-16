package main

import (
	"fmt"
	"strconv"
)

// 型
func main() {
	/* ---------------------------------------------------------------------
	 * 整数型
	 ---------------------------------------------------------------------*/

	var i int = 100

	/*
		最大数　最小数
		int8
		int16
		int32
		int164
	*/
	var i2 int64 = 200
	fmt.Println(i + 50) // 150
	// fmt.Println(i + i2)	// mismatched types

	// 型出力
	fmt.Printf("%T\n", i2) // int64

	// 型変換
	fmt.Printf("%T\n", int32(i2)) // int32

	/* ---------------------------------------------------------------------
	 * 不動小数点型
	 ---------------------------------------------------------------------*/

	var fl64 float64 = 2.4
	fmt.Println(fl64) // 2.4

	fl := 3.2                        // 暗黙敵な宣言の場合、自動で float64に宣言
	fmt.Println(fl64 + fl)           // 5.6
	fmt.Printf("%T, %T\n", fl64, fl) // float64 float64

	var fl32 float32 = 1.2   // 基本的にはあんまり使わない
	fmt.Printf("%T\n", fl32) // float32

	fmt.Printf("%T\n", float64(fl32)) // float64

	// 特殊な値を保持
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN

	/* ---------------------------------------------------------------------
	 * uint(+の整数), complex(複素数)型
	 * ・あまり使用頻度はない
	---------------------------------------------------------------------*/

	var u8 uint = 255
	var c64 complex64 = -5 + 12i
	fmt.Println(u8, c64) // 255 -5+12i

	/* ---------------------------------------------------------------------
	 * bool(論理値)型
	---------------------------------------------------------------------*/
	var t, f bool = true, false
	fmt.Println(t, f) // true false

	/* ---------------------------------------------------------------------
	 * string(文字列)型
	---------------------------------------------------------------------*/
	var s string = "Hello Golang"
	fmt.Println(s)        // Hello Golang
	fmt.Printf("%T\n", s) // string

	// 数字を文字列として扱う
	var si string = "300"
	fmt.Println(si)        // 300
	fmt.Printf("%T\n", si) // string

	// 複数行に渡る文字列を表示 : バッククォーテーション（``）
	/*
	   test
	   test
	   test
	*/
	fmt.Println(`
		test
		test
		test
	`)

	// ダブルクォーテーションを文字列として使いたい場合
	fmt.Println("\"") // "
	fmt.Println(`"`)  // "

	// 文字列は byte配列の集まり
	fmt.Println(s[0])         // 72
	fmt.Println(string(s[0])) // H

	/* ---------------------------------------------------------------------
	 * byte(uint8)型
	---------------------------------------------------------------------*/

	byteA := []byte{72, 73}
	fmt.Println(byteA) // [72 73]

	// byte -> string
	fmt.Println(string(byteA)) // HI

	// string -> byte
	c := []byte("HI")
	fmt.Println(c) // [72 73]

	/* ---------------------------------------------------------------------
	 * 配列型
	 * - size変更ができない
	---------------------------------------------------------------------*/
	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2) // [a b ] = [a b 空文字]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3) // [1, 2, 3]

	// 要素数の省略 = 要素の数を自動でカウント
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)        // [C D]
	fmt.Printf("%T\n", arr4) // [2]string

	// 値の取り出し
	fmt.Println(arr2[0]) // a
	fmt.Println(arr2[1]) // b
	fmt.Println(arr2[2]) // 空文字

	// 値の更新
	arr2[2] = "c"
	fmt.Println(arr2) // [a b c]

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5) //cannot use arr1 (variable of type [3]int) as [4]int value in assignment

	// 配列の要素数を調べる
	fmt.Println(len(arr1)) // 3

	/* ---------------------------------------------------------------------
	 * interface & nil 型
	---------------------------------------------------------------------*/

	// interface : あらゆる型と互換性がある、計算や演算不可
	var x interface{}
	fmt.Println(x) // <nil>：値を何も持っていない状態

	x = 1
	fmt.Println(x) // 1
	x = 3.14
	fmt.Println(x) // 3.14
	x = "str"
	fmt.Println(x) // str
	x = [3]int{1, 2, 3}
	fmt.Println(x) // [1, 2, 3]

	/* ---------------------------------------------------------------------
	 * 型変換
	---------------------------------------------------------------------*/

	// 数値型の相互変換
	var i5 int = 1
	flo64 := float64(i5)
	fmt.Println(flo64)        // 1
	fmt.Printf("%T\n", i5)    // int
	fmt.Printf("%T\n", flo64) // float64

	i6 := int(flo64)
	fmt.Printf("%T\n", i6) // int

	fl2 := 10.5
	i7 := int(fl2)
	fmt.Println(i7)        // 10、小数点以下は切り捨て
	fmt.Printf("%T\n", i7) // int

	// 文字列型　→　数字型
	var s3 string = "100"
	fmt.Printf("%T\n", s3)    // string
	i8, _ := strconv.Atoi(s3) // _ : 関数からの戻り値を破棄
	fmt.Println(i8)           // 8
	fmt.Printf("%T\n", i8)    // int

	// 数字型 → 文字列型
	var i9 int = 200
	s4 := strconv.Itoa(i9)
	fmt.Println(s4)        // 200
	fmt.Printf("%T\n", s4) // string

	// 文字列型　↔　byte配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b) // [72 101 108 108 111 32 87 111 114 108 100]

	h2 := string(b)
	fmt.Println(h2) // Hello World
}
