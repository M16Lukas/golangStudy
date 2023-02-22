package main

import (
	"fmt"
	"time"
)

func main() {

	/* ------------------------------------------------------
	* tiem
	--------------------------------------------------------*/
	// 時間の生成
	// 今の時間
	t := time.Now()
	fmt.Println(t) // 2023-02-22 20:32:45.2459865 +0900 KST m=+0.002158301

	// 指定した時間を生成
	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)              // 2020-06-10 10:10:10 +0900 KST
	fmt.Println(t2.Year())       // 2020
	fmt.Println(t2.YearDay())    // 162
	fmt.Println(t2.Month())      // June
	fmt.Println(t2.Weekday())    // Wednesday
	fmt.Println(t2.Day())        // 10
	fmt.Println(t2.Hour())       // 10
	fmt.Println(t2.Minute())     // 10
	fmt.Println(t2.Second())     // 10
	fmt.Println(t2.Nanosecond()) // 0
	fmt.Println(t2.Zone())       // KST 32400

	// 時刻の感覚を表現する
	// time.Duration型を返す
	fmt.Println(time.Hour)        // 1h0m0s
	fmt.Printf("%T\n", time.Hour) // time.Duration
	fmt.Println(time.Minute)      // 1m0s
	fmt.Println(time.Second)      // 1s
	fmt.Println(time.Millisecond) // 1ms
	fmt.Println(time.Microsecond) // 1µs
	fmt.Println(time.Nanosecond)  // 1ns

	// time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d) // 2h30m0s

	// 現在時刻の2分30秒後を表すtime.Time型の取得
	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3) // 2023-02-22 20:42:53.9994618 +0900 KST m=+135.022612201

	// 時刻の差分を取得
	t4 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	t5 := time.Now()
	fmt.Println(t5) // 2023-02-22 20:44:50.1695931 +0900 KST m=+0.024234501

	// t5 - t6
	d2 := t4.Sub(t5)
	fmt.Println(d2) // -13892h44m50.1695931s

	// 時刻を比較する
	fmt.Println(t5.Before(t4)) // false
	fmt.Println(t5.After(t4))  // true
	fmt.Println(t4.Before(t5)) // true
	fmt.Println(t4.After(t5))  // false

	// 指定時間のスリープ
	// ５秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("５秒停止後表示")

}
