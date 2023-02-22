package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
----------------------------------------------------------------
* bufio
* Goの基本的な入出力処理にバッファ処理を追加した機能をまとめたパッケージ
----------------------------------------------------------------
*/

func main() {
	// 標準入力を行単位で読み込む

	// 標準入力をソースにしたスキャナの生成
	// 改行を区切りとしてスキャン処理する
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャンが成功する限り繰り返すループ
	for scanner.Scan() {
		// スキャン内容を文字列で出力
		fmt.Println(scanner.Text())
	}

	// スキャンにエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み取りエラー", err)
	}

}
