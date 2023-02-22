package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/* ------------------------------------------------------
	* os
	------------------------------------------------------*/

	// exit
	os.Exit(1)
	fmt.Println("start") // 実行されない

	defer func() {
		fmt.Println("defer")
	}() // 実行されない
	os.Exit(0)

	// log.Fatal
	// エラーが出力されて、os.Exit()が実行される
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Args
	// ビルド：go build -o getargs
	// 実行： ./getargs 123 456 789
	fmt.Println(os.Args[0]) // ./getargs
	fmt.Println(os.Args[1]) // 123
	fmt.Println(os.Args[2]) // 456
	fmt.Println(os.Args[3]) // 789

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args)) // length=4
	// os.Argsの中身をすべて表示
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	// ファイル操作
	// open
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // リソースの解放

	// Create
	f2, _ := os.Create("foo.txt") // ファイルを作成
	f2.Write([]byte("Hello\n"))
	f2.WriteAt([]byte("Golang"), 6)
	f2.Seek(0, os.SEEK_END) // SEEK_END : ファイルの末尾にオフセットを移動
	f2.WriteString("Yeah")

	// Read
	f3, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f3.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs) // n : 読みとんだバイト数
	fmt.Println(n)       // 16
	fmt.Println(string(bs))
	// Hello
	// GolangYeah

	bs2 := make([]byte, 128)
	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)          // 6
	fmt.Println(string(bs2)) // ngYeah

	// OpenFile
	// O_RDONLY 読み込み専用
	// O_WRONLY 書き込み専用
	// O_RDWR   読み書き可能
	// O_APPEND ファイル末尾に追記
	// O_CREATE ファイルがなければ作成
	// O_TRUNC  可能であればファイルの内容をオープン時に空にする
	f4, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f4.Close()

	bs3 := make([]byte, 128)
	n2, err := f4.Read(bs3)
	fmt.Println(n2) // 16
	fmt.Println(string(bs3))
	// Hello
	// GolangYeah
}
