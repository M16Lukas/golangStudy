package main

import (
	"fmt"
	"sync"
	"time"
)

/*
----------------------------------------
* sync
* ミューテックスによる同期処理
----------------------------------------
*/

var st2 struct{ A, B, C int }

// ミューテックスを保持するパッケージ変数
// 任意の処理のブロックを同期化するために利用
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ロック
	// ロックされている間は、一つのゴルーチンしか処理を行えない
	mutex.Lock()

	st2.A = n
	time.Sleep(time.Microsecond)
	st2.B = n
	time.Sleep(time.Microsecond)
	st2.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st2)

	// アンロック
	mutex.Unlock()
}

func main() {
	// ミューテックスを生成
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {
	}
}
