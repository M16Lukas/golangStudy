package main

/*
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
*/

import (
	"context"
	"fmt"
	"time"
)

/*
---------------------------------------------
* context
* APIのサーバーやクライアントを使う時に
* contextを提供してキャンセルやタイムアウトをできる仕組み
---------------------------------------------
*/

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	time.Sleep(2 * time.Second)
	fmt.Println("終了")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)
	// contextの作成
	ctx := context.Background()
	// タイムアウトを付けて再定義
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// リソースの解放
	defer cancel()

	go longProcess(ctx, ch)

	//cancel()

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("##########Error###########")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}

	}

	fmt.Println("ループ抜けた")
}
