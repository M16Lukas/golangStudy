package main

import "fmt"

func main() {

	/* -----------------------------------------------------------------------
	* 関数
	-----------------------------------------------------------------------*/
	i := Plus(1, 2)
	fmt.Println(i) // 3

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3) // 3 0

	// 戻り値の破棄
	i4, _ := Div(7, 4)
	fmt.Println(i4) // 1

	i5 := Double(1000)
	fmt.Println(i5) // 2000

	Noreturn() // No Return

	/* -----------------------------------------------------------------------
	* 無名関数
	-----------------------------------------------------------------------*/
	f := func(x, y int) int {
		return x + y
	}
	f2 := f(1, 2)
	fmt.Println(f2) // 3

	f3 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(f3) // 3

	/* -----------------------------------------------------------------------
	* 関数を返す関数
	-----------------------------------------------------------------------*/
	f4 := ReturnFunc()
	f4() // I'm a Function

	/* -----------------------------------------------------------------------
	* 関数を引数に取る関数
	-----------------------------------------------------------------------*/
	CallFunc(func() {
		fmt.Println("I'm a function")
	}) // I'm a function

	/* -----------------------------------------------------------------------
	* クロージャー
	* 関数と関数の処理に関する関数外の環境をセットして閉じ込めた物
	-----------------------------------------------------------------------*/
	f5 := Later()
	fmt.Println(f5("Hello"))  //
	fmt.Println(f5("my"))     // Hello
	fmt.Println(f5("name"))   // my
	fmt.Println(f5("is"))     // name
	fmt.Println(f5("golang")) // is

	/* -----------------------------------------------------------------------
	* ジェネレーター
	* 何らかのルールに従って連続した値を返し続ける仕組み
	-----------------------------------------------------------------------*/
	ints := integers()
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3
	fmt.Println(ints()) // 4

	otherints := integers()
	fmt.Println(otherints()) // 1
	fmt.Println(otherints()) // 2
	fmt.Println(otherints()) // 3
}

/* -----------------------------------------------------------------------
* 関数
-----------------------------------------------------------------------*/

func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y

	return q, r
}

// 返り値を表す
func Double(price int) (result int) {
	result = price * 2
	return
}

// 返り値がない関数
func Noreturn() {
	fmt.Println("No Return")
	return
}

/* -----------------------------------------------------------------------
* 関数を返す関数
-----------------------------------------------------------------------*/
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a Function")
	}
}

/* -----------------------------------------------------------------------
* 関数を引数に取る関数
-----------------------------------------------------------------------*/
func CallFunc(f func()) {
	f()
}

/* -----------------------------------------------------------------------
* クロージャー
* 関数と関数の処理に関する関数外の環境をセットして閉じ込めた物
-----------------------------------------------------------------------*/
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

/* -----------------------------------------------------------------------
* ジェネレーター
* 何らかのルールに従って連続した値を返し続ける仕組み
-----------------------------------------------------------------------*/
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}