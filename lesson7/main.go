package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
---------------------------------------------------------
* switch型スイッチ
---------------------------------------------------------
*/
func anything(a interface{}) {
	fmt.Println(a)
}

/*
---------------------------------------------------------
* defer
* 関数の終了時に実行される処理を登録
---------------------------------------------------------
*/
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

/*
---------------------------------------------------------
* go goroutin
* goの並行処理
---------------------------------------------------------
*/
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

/*
---------------------------------------------------------
* init
* init関数はmain関数より先に呼ばれる
---------------------------------------------------------
*/
func init() {
	fmt.Println("init")
}

/* ---------------------------------------------------------
 * 制御構文
---------------------------------------------------------*/

func main() {

	/* ---------------------------------------------------------
	 	* if
	---------------------------------------------------------*/
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	} // I don't know

	// 簡易文付き if
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	} // one hundred

	x := 0
	if x := 2; true {
		fmt.Println(x)
	} // 2
	fmt.Println(x) // 0

	/* ---------------------------------------------------------
	 	* エラーハンドリング
	---------------------------------------------------------*/
	var s string = "100"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", i) // int

	/* ---------------------------------------------------------
	 	* for
	---------------------------------------------------------*/

	// 無限ループ
	// for {
	// 	fmt.Println("Loop")
	// }

	// 条件付き for
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// 一般的な for
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // 以下の処理をスキップして、for文に戻る
		}
		if i == 6 {
			break // ループを強制終了
		}
		fmt.Println(i)
	}

	// 配列処理
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// range
	// インデックス、値
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 値だけ
	for _, v := range arr {
		fmt.Println(v)
	}

	// インデックスだけ
	for i := range arr {
		fmt.Println(i)
	}

	// スライス
	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// map[key]value
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	/* ---------------------------------------------------------
	 	* switch
	---------------------------------------------------------*/
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	} // 1 or 2

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I dont know")
	}

	n1 := 6
	switch {
	case n1 > 0 && n1 < 4:
		fmt.Println("0 < n < 4")
	case n1 > 3 && n1 < 7:
		fmt.Println("3 < n < 7")
	}

	/* ---------------------------------------------------------
	* switch型スイッチ
	---------------------------------------------------------*/
	anything("Aaa") // Aaa
	anything(1)     // 1

	var x2 interface{} = 3
	i, isInt := x2.(int)  // interface型　→　int型
	fmt.Println(i, isInt) // 3 true
	// fmt.Println(x2 + 2) // error

	f, isFloat64 := x2.(float64)
	fmt.Println(f, isFloat64) // 0 false

	if x2 == nil {
		fmt.Println("None")
	} else if i, isInt := x2.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x2.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I dont know")
	} // 3 x is Int

	//　型によるスイッチ文
	switch x2.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I dont know")
	} // int

	switch v := x2.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	} // 3 int

	/* ---------------------------------------------------------
	* ラベル付きfor
	---------------------------------------------------------*/
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("start")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理をしない")
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END")

	// Loop:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 3; j++ {
	// 			if j > 1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(i, j, i*j)
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}

	/* ---------------------------------------------------------
	* defer
	* 関数の終了時に実行される処理を登録
	---------------------------------------------------------*/
	TestDefer()
	// START
	// END

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	// 1
	// 2
	// 3

	// 注意：後から登録された式から表示
	RunDefer()
	// 3
	// 2
	// 1

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))

	/* ---------------------------------------------------------
	* panic & recover
	* 例外処理
	---------------------------------------------------------*/
	// recover : panic状態であれば、xに値が帰ってくる
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// panic : 実行中のプログラムを強制的に終了
	panic("runtime error") // panic runtime error
	fmt.Println("START")

	/* ---------------------------------------------------------
	* go goroutin
	* goの並行処理
	---------------------------------------------------------*/
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond)
	}

	/* ---------------------------------------------------------
	* init
	* init関数はmain関数より先に呼ばれる
	---------------------------------------------------------*/
	fmt.Println("Main")
	// init
	// Main
}
