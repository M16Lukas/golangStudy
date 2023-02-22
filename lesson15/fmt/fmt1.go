package main

import (
	"fmt"
	"os"
)

func main() {

	/* -----------------------------------------------------
	* fmt
	-----------------------------------------------------*/

	fmt.Println("表示") // 表示

	// fmt標準、改行なし
	print("Hello")
	// 改行
	println("Hello!") // HelloHello!

	fmt.Print("Hello")
	fmt.Println("Hello!") // HelloHello!

	// Println : 各々の文字列は半角スペースで区切られ、文字列の最後に改行を追加して連結して表示
	fmt.Println("Hello", "world!") // Hello world!
	fmt.Println("Hello", "world!") // Hello world!

	// Printf : フォーマットを指定
	fmt.Printf("%s\n", "Hello")  // Hello
	fmt.Printf("%#v\n", "Hello") // "Hello", #v : 値をそのまま表示

	// Sprint系：出力ではなく、フォーマットした結果を文字列で返す。
	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", "Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s) // Hello
	fmt.Println(s1)
	/*
	 Hello

	*/
	fmt.Println(s2)
	/*
	 Hello

	*/

	// Fprint系：書き込み先指定、任意のI/Oライター型への出力関数
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")
	// HelloHelloHello

	f, _ := os.Create("test.txt")
	defer f.Close()
	fmt.Fprintln(f, "Fprint")
}
