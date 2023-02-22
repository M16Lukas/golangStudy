package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
--------------------------------------------------------------
* ioutil
* 入出力処理をサポートする機能がまとめられたパッケージ
* 簡易的なファイルの読み書きを用意している
* os パッケージでも対応することができるが、目的によって使い分けられる
--------------------------------------------------------------
*/

func main() {
	// 入力全体を読み込む
	f, _ := os.Open("doc.txt")
	bs, _ := ioutil.ReadAll(f) // 巨大な入力データに使うのは適していないことに注意
	fmt.Println(string(bs))

	// ファイルの読み込み
	bs2, _ := ioutil.ReadFile("doc.txt")
	fmt.Println(string(bs2))

	// ファイルの書き込み
	if err := ioutil.WriteFile("bar.txt", bs2, 0666); err != nil {
		log.Fatalln(err)
	}

}
