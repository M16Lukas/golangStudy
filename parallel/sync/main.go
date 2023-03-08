package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

type Person struct {
	Name string
}

func Hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v\n", id)
}

func main() {
	/*
	* WaitGroup
	* 多数のゴルーチンを実行しているジョブの終了待ちに使用
	 */
	// var wg sync.WaitGroup

	// wg.Add(1) // +1
	// go func() {
	// 	defer wg.Done() // -1
	// 	fmt.Println("1st Droutine Start")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("1st Droutine Done")
	// }()

	// wg.Add(1) // +1
	// go func() {
	// 	defer wg.Done() // -1
	// 	fmt.Println("2st Droutine Start")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("2st Droutine Done")
	// }()

	// var CPU int = runtime.NumCPU()
	// wg.Add(CPU)
	// for i := 0; i < CPU; i++ {
	// 	go Hello(&wg, i)
	// }

	// wg.Wait() // 追加したゴルーチンの数が０になるまで待ち

	/*
	* 競合 Mutex
	 */

	// var memoryAccess sync.Mutex
	// var data int

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	memoryAccess.Lock()
	// 	data++
	// 	memoryAccess.Unlock()
	// }()

	// wg.Wait()

	// memoryAccess.Lock()
	// if data == 0 {
	// 	fmt.Println(0)
	// } else {
	// 	fmt.Println(data)
	// }
	// memoryAccess.Unlock()

	/*
	* デッドロック
	 */
	// printSum := func(v1, v2 *value) {
	// 	defer wg.Done()

	// 	v1.mu.Lock()
	// 	fmt.Printf("%v がロックを取得しました\n", v1.name)
	// 	defer v1.mu.Unlock()

	// 	time.Sleep(2 * time.Second)

	// 	v2.mu.Lock()
	// 	fmt.Printf("%v がロックを取得しました\n", v2.name)
	// 	defer v2.mu.Unlock()

	// 	fmt.Println(v1.value + v2.value)
	// }

	// var a value = value{
	// 	name: "a",
	// }

	// var b value = value{
	// 	name: "b",
	// }

	// wg.Add(2)

	// go printSum(&a, &b)
	// go printSum(&b, &a)

	// wg.Wait()

	/*
	* リソース枯渇
	 */
	// var lock sync.Mutex
	// const timer = 1 * time.Second

	// // 貪欲なワーカー
	// greedyWorker := func() {
	// 	defer wg.Done()

	// 	count := 0

	// 	begin := time.Now()
	// 	for time.Since(begin) <= timer {
	// 		lock.Lock()
	// 		time.Sleep(3 * time.Nanosecond)
	// 		lock.Unlock()
	// 		count++
	// 	}

	// 	fmt.Printf("greedyWorker : %v\n", count)
	// }

	// // 丁寧なワーカー
	// politeWorker := func() {
	// 	defer wg.Done()

	// 	count := 0

	// 	begin := time.Now()
	// 	for time.Since(begin) <= timer {
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		count++
	// 	}

	// 	fmt.Printf("politeWorker : %v\n", count)
	// }

	// wg.Add(2)
	// go greedyWorker()
	// go politeWorker()
	// wg.Wait()

	/*
	*  MutexとRWMutex
	* RWMutex : 読み込みのみ許可するロックが可能
	 */
	// var count int
	// var lock sync.RWMutex

	// increment := func(wg *sync.WaitGroup, l sync.Locker) {
	// 	l.Lock()
	// 	defer l.Unlock()
	// 	defer wg.Done()

	// 	fmt.Println("increment")
	// 	count++
	// 	time.Sleep(1 * time.Second)
	// }

	// read := func(wg *sync.WaitGroup, l sync.Locker) {
	// 	l.Lock()
	// 	defer l.Unlock()
	// 	defer wg.Done()

	// 	fmt.Println("read")
	// 	fmt.Println(count)
	// 	time.Sleep(1 * time.Second)
	// }

	// start := time.Now()

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go increment(&wg, &lock)
	// }

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go read(&wg, lock.RLocker())
	// }

	// wg.Wait()
	// end := time.Now()

	// fmt.Println(end.Sub(start))

	/*
	*  Cond
	* 条件変数と呼ばれる排他制御の仕組み
	* ロックを掛けたり、解除したりすることでクリティカルセクションを保護する為に
	 */
	// var mutex sync.Mutex
	// cond := sync.NewCond(&mutex)

	// for _, name := range []string{"A", "B", "C"} {
	// 	go func(name string) {
	// 		mutex.Lock()
	// 		defer mutex.Unlock()

	// 		cond.Wait()
	// 		fmt.Println(name)
	// 	}(name)
	// }

	// fmt.Println("Ready....")
	// time.Sleep(time.Second)
	// fmt.Println("Go!")

	// cond.Broadcast()

	// time.Sleep(time.Second)
	// fmt.Println("Done")

	/*
	* ライブロック
	* プログラムの状態を全く進めていないプログラムを指す
	* 並行処理の再開のために永遠に反応しない状態になっている
	 */
	// cond := sync.NewCond(&sync.Mutex{})

	// go func() {
	// 	for range time.Tick(1 * time.Second) {
	// 		cond.Broadcast()
	// 	}
	// }()

	// var flag [2]bool

	// takeStep := func() {
	// 	cond.L.Lock()
	// 	cond.Wait()
	// 	cond.L.Unlock()
	// }

	// p0 := func() {
	// 	defer wg.Done()
	// 	flag[0] = true
	// 	takeStep()

	// 	for flag[1] {
	// 		takeStep()

	// 		flag[0] = false

	// 		takeStep()

	// 		if flag[0] != flag[1] {
	// 			break
	// 		}

	// 		takeStep()

	// 		flag[0] = true
	// 	}
	// }

	// p1 := func() {
	// 	defer wg.Done()
	// 	flag[1] = true
	// 	takeStep()

	// 	for flag[0] {
	// 		takeStep()

	// 		flag[1] = false

	// 		takeStep()

	// 		if flag[0] != flag[1] {
	// 			break
	// 		}

	// 		takeStep()

	// 		flag[1] = true
	// 	}
	// }

	// wg.Add(2)
	// go p0()
	// go p1()

	// wg.Wait()

	/*
	* Once
	* Doが呼び出された回数をカウント
	 */
	// count := 0

	// increment := func() {
	// 	count++
	// }

	// decrement := func() {
	// 	count--
	// }

	// var once sync.Once

	// once.Do(increment)
	// once.Do(decrement)

	// fmt.Println(count) // 1

	/*
	* Pool
	* オブジェクトプールパターンを並行処理で安全な形で実装したもの
	* オブジェクトプールパターン : 使うものを決まった数だけ作るという方法
	* たくさんのリクエストがあった際に、このプールしたオブジェクトの中から使う
	 */
	// mypool := &sync.Pool{
	// 	New: func() interface{} {
	// 		fmt.Println("Create new instance")
	// 		return new(Person)
	// 	},
	// }

	// mypool.Put(&Person{Name: "1"})
	// mypool.Put(&Person{Name: "2"})

	// instance1 := mypool.Get()
	// instance2 := mypool.Get()
	// instance3 := mypool.Get()

	// fmt.Println(instance1, instance2, instance3) // &{1} &{2} &{}

	/*
	* Map
	 */
	// smap := &sync.Map{}
	// // データの貯蔵
	// smap.Store("Hello", "World")
	// smap.Store(1, 2)

	// // ループ
	// smap.Range(func(key, value interface{}) bool {
	// 	fmt.Println(key, value)
	// 	return true
	// })

	// // データの削除
	// smap.Delete(1)

	// // データの取得
	// v, ok := smap.Load("Hello")
	// if ok {
	// 	fmt.Println(v)
	// }

	// // もしマップに値が存在しない場合、値を追加する
	// smap.LoadOrStore("Hello", "Woooooooooooorld")
	// smap.LoadOrStore("2", "123")

	/*
	* Atomic
	* 不可分操作と呼ばれる操作を提供するパッケージ
	* Mutexと同様な処理を簡単に実行
	 */
	// var count int64

	// increment := func() {
	// 	atomic.AddInt64(&count, 1)
	// }

	// increment()
	// fmt.Println(count)

	/*
	* channel
	* FIFO型のデーター構造
	* 順序よく受け渡す為のデーター構造
	* 整合性が壊れる心配がない
	 */

	// 同期
	// 読み書きが準備できるまでブロックする
	// ch := make(chan string)

	// go func() {
	// 	ch <- "Hello"
	// }()

	// v, ok := <-ch
	// fmt.Println(v, ok) // Hello true

	// // close
	// begin := make(chan interface{})

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	fmt.Printf("Start goroutine %d\n", i)

	// 	go func(i int) {
	// 		defer wg.Done()
	// 		<-begin
	// 		fmt.Printf("%d has begin\n", i)
	// 	}(i)
	// }

	// fmt.Println("Unblocking goroutine")
	// close(begin)
	// wg.Wait()

	// // buffer
	// ch1 := make(chan int, 5)

	// go func() {
	// 	defer fmt.Println("Close")
	// 	defer close(ch1)
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Printf("Writing to channel: %v\n", i)
	// 		ch1 <- 1
	// 	}
	// }()

	// for integer := range ch1 {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Printf("reading to channel: %v\n", integer)
	// }
	/*
		Writing to channel: 0
		Writing to channel: 1
		Writing to channel: 2
		Writing to channel: 3
		Writing to channel: 4
		Close
		reading to channel: 1
		reading to channel: 1
		reading to channel: 1
		reading to channel: 1
		reading to channel: 1
	*/

	// ライフサイクルのカプセル化

	// 読み取り専用
	// chanOwner := func() <-chan int {
	// 	// 初期化
	// 	resultStream := make(chan int, 5)

	// 	go func() {
	// 		defer close(resultStream)
	// 		for i := 0; i < 5; i++ {
	// 			resultStream <- i
	// 		}
	// 	}()

	// 	return resultStream
	// }

	// resultStream := chanOwner()
	// for result := range resultStream {
	// 	fmt.Printf("Received: %v\n", result)
	// }

	// fmt.Println("Done")

	/*
		Received: 0
		Received: 1
		Received: 2
		Received: 3
		Received: 4
		Done
	*/

	/*
	* select for-select & timeout & cancel & default
	* select : 複数のチャネルの操作をまとめる
	 */
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch2 <- i
		}
	}()

loop:
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				break loop
			}
			fmt.Printf("ch1 : %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				break loop
			}
			fmt.Printf("ch2 : %v\n", v)
		default:
			fmt.Println("default")
		}
	}
}
