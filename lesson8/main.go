package main

import (
	"fmt"
	"time"
)

/*
* 参照型
 */

func main() {
	/* -----------------------------------------------------------
	 * スライス
	 ------------------------------------------------------------*/
	var sl []int
	fmt.Println(sl) // []

	// 明示的宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // [100, 200]

	// 暗黙的宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3) // [A B]

	sl4 := make([]int, 5)
	fmt.Println(sl4) // [0 0 0 0 0]

	// 値の更新
	sl2[0] = 1000
	fmt.Println(sl2) // 1000 200

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)                 // [1 2 3 4 5]
	fmt.Println(sl5[0])              // 1
	fmt.Println(sl5[2:4])            // [3 4]
	fmt.Println(sl5[:2])             // [1 2]
	fmt.Println(sl5[2:])             // [3 4 5]
	fmt.Println(sl5[:])              // [1 2 3 4 5]
	fmt.Println(len(sl5) - 1)        // 5
	fmt.Println(sl5[1 : len(sl5)-1]) // [2 3 4]

	/* -----------------------------------------------------------
	 * スライス append make len cap
	 ------------------------------------------------------------*/
	// append
	sl6 := []int{100, 200}
	sl6 = append(sl6, 300)
	fmt.Println(sl6) // [100 200 300]
	sl6 = append(sl6, 400, 500, 600)
	fmt.Println(sl6) // [100 200 300 400 500 600]

	// make
	sl7 := make([]int, 5)
	fmt.Println(sl7) // [0 0 0 0 0]

	// len
	fmt.Println(len(sl7)) // 5

	// cap(容量)
	fmt.Println(cap(sl7)) // 5

	sl8 := make([]int, 5, 10)
	fmt.Println(len(sl8)) // 5
	fmt.Println(cap(sl8)) // 10

	// 要領以上の要素が追加されるとメモリの消費が倍になる
	sl8 = append(sl8, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl8)) // 12
	fmt.Println(cap(sl8)) // 20

	/* -----------------------------------------------------------
	 * スライス copy
	 ------------------------------------------------------------*/
	sl9 := []int{1, 2, 3, 4, 5}
	sl10 := make([]int, 5, 10)
	fmt.Println(sl10)    // [0 0 0 0 0]
	n := copy(sl10, sl9) // n : copyに成功した数
	fmt.Println(n, sl10) // 5 [1 2 3 4 5]

	/* -----------------------------------------------------------
	 * スライス for
	 ------------------------------------------------------------*/
	sl11 := []string{"A", "B", "C"}
	fmt.Println(sl11) // [A B C]

	for i, v := range sl11 {
		fmt.Println(i, v)
	}

	for i := 0; i < len(sl11); i++ {
		fmt.Println(sl11[i])
	}

	/* -----------------------------------------------------------
	 * スライス 可変長引数
	 ------------------------------------------------------------*/
	fmt.Println(Sum())                              // 0
	fmt.Println(Sum(1, 2, 3))                       // 6
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55

	sl12 := []int{1, 2, 3}
	fmt.Println(Sum(sl12...)) // 6

	/* -----------------------------------------------------------
	 * マップ
	 ------------------------------------------------------------*/
	// 明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m) // map[A:100 B:200]

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2) // map[A:100 B:200]

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3) // map[1:A 2:B]

	// make関数
	m4 := make(map[int]string)
	fmt.Println(m4) // map[]

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4) // map[1:JAPAN 2:USA]

	// 値の取り出し
	fmt.Println(m["A"]) // 100
	fmt.Println(m4[2])  // USA
	fmt.Println(m4[3])  // 空文字

	s, ok := m4[1]
	fmt.Println(s, ok) // JAPAN true
	s2, ok := m4[3]
	fmt.Println(s2, ok) //   false

	// 値の更新
	m4[2] = "US"
	fmt.Println(m4) // map[1:JAPAN 2:US]
	m4[3] = "CHINA"
	fmt.Println(m4) // map[1:JAPAN 2:US 3:CHINA]

	// 値の削除
	delete(m4, 3)
	fmt.Println(m4) // map[1:JAPAN 2:US]

	// len
	fmt.Println(len(m4)) // 2

	/* -----------------------------------------------------------
	 * マップ for
	 ------------------------------------------------------------*/
	m5 := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m5 {
		fmt.Println(k, v)
	}

	/* -----------------------------------------------------------
	 * チャンネル(channel)
	 * 複数のゴルーチン間でデータを受け渡しをするために設計されたデータ構造
	 * 宣言、操作
	 * FIFO性質
	 ------------------------------------------------------------*/

	// 受信専用のチャンネル
	// var ch2 <-chan int

	// 送信専用のチャンネル
	// var ch3 chan<- int

	var ch1 chan int
	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1)) // 0
	fmt.Println(cap(ch2)) // 0

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // 5

	// データの操作
	// チャンネルにデータを送信
	ch3 <- 1
	fmt.Println(len(ch3)) // 1

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3)) // 3

	// チャンネルからデータを受信
	i := <-ch3
	fmt.Println(i)        // 1
	fmt.Println(len(ch3)) // 2

	i2 := <-ch3
	fmt.Println(i2)       // 2
	fmt.Println(len(ch3)) // 1

	fmt.Println(<-ch3)    // 3
	fmt.Println(len(ch3)) // 0

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	// ch3 <- 6	// error : deadlock

	/* -----------------------------------------------------------
	 * チャンネルとゴルーチン
	 ------------------------------------------------------------*/
	ch4 := make(chan int)
	ch5 := make(chan int)
	// fmt.Println(<-ch4)	// error : deadlock
	go reciever(ch4)
	go reciever(ch5)

	i3 := 0
	for 2 < 100 {
		ch4 <- i3
		ch5 <- i3
		time.Sleep(50 * time.Millisecond)
		i3++
	}

	/* ----------------------------------------------------------
	 * チャンネル close
	 * closeされたチャンネルに送信：X
	 * closeされたチャンネルから受信：O
	 ------------------------------------------------------------*/
	ch6 := make(chan int, 2) // open状態

	// ch6 <- 1

	// close(ch6) // open　→　close

	// i4, ok := <-ch6
	// fmt.Println(i4, ok) // 1 true
	// i5, ok := <-ch6
	// fmt.Println(i5, ok) // 0 false

	// ゴルーチンとcloseの実例
	go reciever2("1.goroutin", ch6)
	go reciever2("2.goroutin", ch6)
	go reciever2("3.goroutin", ch6)

	i6 := 0
	for i6 < 100 {
		ch6 <- i6
		i6++
	}
	close(ch6)
	time.Sleep(3 * time.Second) // 3sec

	/* -----------------------------------------------------------
	 * チャンネル for
	 ------------------------------------------------------------*/
	ch7 := make(chan int, 3)
	ch7 <- 1
	ch7 <- 2
	ch7 <- 3
	close(ch7) // deadlock回避：for文を使う時はfor文の前に必ず「close」する

	for i := range ch7 {
		fmt.Println(i)
	}

	/* -----------------------------------------------------------
	 * チャンネル select
	 * 複数チャンネルに対する送受信をゴルーチンを定義させることなく制御できる
	 ------------------------------------------------------------*/
	ch8 := make(chan int, 2)
	ch9 := make(chan string, 2)

	ch9 <- "A"
	ch8 <- 1
	ch9 <- "B"
	ch8 <- 2

	// case実行はランダム：switchとの違い
	select {
	case v1 := <-ch8:
		fmt.Println(v1 + 1000)
	case v2 := <-ch9:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	// 活用例
	ch10 := make(chan int)
	ch11 := make(chan int)
	ch12 := make(chan int)

	// reciever
	go func() {
		for {
			i := <-ch10
			ch11 <- i * 2
		}
	}()

	go func() {
		i2 := <-ch11
		ch12 <- i2 - 1
	}()

	n2 := 0
L:
	for {
		select {
		case ch10 <- n2:
			n2++
		case i3 := <-ch12:
			fmt.Println("recieved", i3)
		default:
			if n > 100 {
				break L
			}
		}
	}
}

/*
-----------------------------------------------------------
* スライス 可変長引数
------------------------------------------------------------
*/
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

/*
-----------------------------------------------------------
* チャンネルとゴルーチン
------------------------------------------------------------
*/
func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

/*
-----------------------------------------------------------
* チャンネル close
------------------------------------------------------------
*/
func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}
