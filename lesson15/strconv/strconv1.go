package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* --------------------------------------------------
	* strconv
	* goの基本的なデータとstring型の相互変換をサポート
	--------------------------------------------------*/

	// 真偽値　→　文字列
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt)) // string

	// 整数　→　文字列
	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", i, i) // -100, string

	// 簡易的に変換できる
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2) // 100, string

	// 浮動小数点型に変換する
	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64)) // 1.23456E+02
	// 指数表現による文字列化（小数点以下２桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64)) // 1.23e+02
	// 実数表現による文字列化
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64)) // 123.456
	//  実数表現による文字列化（小数点下２桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64)) // 123.46
	// 指数部の大きさで変動する表現による文字列化
	fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))       // 123.456
	fmt.Println(strconv.FormatFloat(123456789.123, 'f', -1, 64)) // 123456789.123
	// 指数部の大きさで変動する表現による文字列化（仮数部全体が４桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64)) // 123.5
	// 指数部の大きさで変動する表現による文字列化（仮数部全体が８桁まで）
	fmt.Println(strconv.FormatFloat(123456789.123, 'G', 8, 64)) // 1.2345679E+08

	// 文字列　→　真偽値
	// trueへ変換できる文字列
	bt1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", bt1, bt1) // true, bool
	bt2, _ := strconv.ParseBool("1")
	bt3, _ := strconv.ParseBool("t")
	bt4, _ := strconv.ParseBool("T")
	bt5, _ := strconv.ParseBool("TRUE")
	bt6, _ := strconv.ParseBool("True")
	fmt.Println(bt2, bt3, bt4, bt5, bt6) // true true true true true

	// falseへ変換できる文字列
	// 二番目の戻り値はerror型なので、エラーハンドリングも可能
	bf1, ok := strconv.ParseBool("false")
	if ok != nil {
		fmt.Println("Convert Error")
	}
	fmt.Printf("%v, %T\n", bf1, bf1) // false, bool
	bf2, _ := strconv.ParseBool("0")
	bf3, _ := strconv.ParseBool("f")
	bf4, _ := strconv.ParseBool("F")
	bf5, _ := strconv.ParseBool("FALSE")
	bf6, _ := strconv.ParseBool("False")
	fmt.Println(bf2, bf3, bf4, bf5, bf6) // false false false false false

	// 文字列　→　整数型
	i3, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("%v, %T\n", i3, i3) // 12345, int64
	i4, _ := strconv.ParseInt("-1", 10, 0)
	fmt.Printf("%v, %T\n", i4, i4) // -1, int64

	// 簡易的に変換
	i10, _ := strconv.Atoi("123")
	fmt.Println(i10) // 123

	// 文字列　→　浮動小数点型
	fl1, _ := strconv.ParseFloat("3.14", 64)
	fl2, _ := strconv.ParseFloat(".2", 64)
	fl3, _ := strconv.ParseFloat("-2", 64)
	fl4, _ := strconv.ParseFloat("1.2345e8", 64)
	fl5, _ := strconv.ParseFloat("1.2345E8", 64)
	fmt.Println(fl1, fl2, fl3, fl4, fl5) // 3.14 0.2 -2 1.2345e+08 1.2345e+08

}
