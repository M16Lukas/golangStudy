package main

import "fmt"

func main() {
	/* ------------------------------------------------------
	 * 算術演算子
	------------------------------------------------------*/
	// 足し算
	fmt.Println(1 + 2)         //3
	fmt.Println("ABC" + "CDE") // ABCCDE
	// 引き算
	fmt.Println(5 - 1) //4
	// 掛け算
	fmt.Println(5 * 4) //20
	// 割り算
	fmt.Println(8 / 4) //2
	// あまりの計算
	fmt.Println(9 % 4) //1

	// 算術演算子の代入
	n := 0
	n += 4         // n = n + 4
	fmt.Println(n) // 4
	n++            // n = n + 1
	fmt.Println(n) // 5
	n--            // n = n - 1
	fmt.Println(n) // 4

	s := "ABC"
	s += "DEF"
	fmt.Println(s) // ABCDEF

	/* ------------------------------------------------------
	 * 比較演算子
	------------------------------------------------------*/
	fmt.Println(1 == 1)        // true
	fmt.Println(1 == 2)        // false
	fmt.Println(4 <= 8)        // true
	fmt.Println(1 >= 8)        // false
	fmt.Println(1 < 8)         // true
	fmt.Println(1 > 8)         // false
	fmt.Println(true == false) // false
	fmt.Println(true != false) // true

	/* ------------------------------------------------------
	 * 論理演算子
	------------------------------------------------------*/
	fmt.Println(true && false == true)  // false
	fmt.Println(true && true == true)   // true
	fmt.Println(true || true == false)  // true
	fmt.Println(false || true == false) // false
	fmt.Println(!true)                  // false
	fmt.Println(!false)                 // true

	/* ------------------------------------------------------
	 * シフト
	------------------------------------------------------*/
	fmt.Println(1 << 0) // 0001 0001（= 1）
	fmt.Println(1 << 1) // 0001 0010（= 2）
	fmt.Println(1 << 2) // 0001 0100（= 4）
	fmt.Println(1 << 3) // 0001 1000（= 8）
}
