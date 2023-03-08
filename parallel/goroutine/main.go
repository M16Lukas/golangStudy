package main

import (
	"fmt"
	"runtime"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}

func memConsumed() uint64 {
	runtime.GC()

	var s runtime.MemStats
	runtime.ReadMemStats(&s)

	return s.Sys // OSから取得したメモリの合計バイト数
}

func main() {
	/*
	* Goroutine基礎
	 */
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("HEEEEELO")
	}()

	wg.Wait()

	/*
	* クロージャとforloopを使ったGoroutineの注意点
	 */
	say := "Hello"

	wg.Add(1)
	go func() {
		defer wg.Done()
		say = "Good Bye"
	}()

	wg.Wait()

	// 子のゴルーチンと親のゴルーチンは同じアドレス空間で実行される
	fmt.Println(say) // Good Bye

	tasks := []string{"A", "B", "C"}

	for _, task := range tasks {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(task)
		}()
		// C
		// C
		// C
	}

	wg.Wait()

	/*
	* ゴルーチンのメモリ消費量
	 */
	var ch <-chan interface{}

	noop := func() {
		wg.Done()
		<-ch
	}

	const numGoroutines = 10000

	wg.Add(numGoroutines)

	before := memConsumed()

	for i := 0; i < numGoroutines; i++ {
		go noop()
	}

	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3f kb", float64(after-before)/numGoroutines/1000)
}
