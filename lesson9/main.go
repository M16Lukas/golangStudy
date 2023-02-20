package main

import "fmt"

/*
------------------------------------------------------------
* ポインタ
* 値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報のこと
* 主に値型の操作に使われ、
* 参照型(スライス、マップ、チャンネル)は元からその機能を含んでいる為、基本的には不要
------------------------------------------------------------
*/

func main() {
	var n int = 100
	fmt.Println(n)  // 100
	fmt.Println(&n) // 0xc000016098

	Double(n)
	fmt.Println(n) // 100

	// ポインタ宣言
	var p *int = &n
	fmt.Println(p)  // 0xc000016098
	fmt.Println(*p) // 100

	*p = 300
	fmt.Println(n) // 300
	n = 200
	fmt.Println(*p) // 200

	DoubleV2(&n)
	fmt.Println(n) // 400

	DoubleV2(p)
	fmt.Println(*p) // 800

	var sl []int = []int{1, 2, 3}
	DoubleV3(sl)
	fmt.Println(sl) // [2 4 6]
}

func Double(i int) {
	i = i * 2
}

func DoubleV2(i *int) {
	*i = *i * 2
}

func DoubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}
