package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

/*
-----------------------------------
* crypt MD5ハッシュ
-----------------------------------
*/

func main() {
	// MD5ハッシュ値を生成
	// 任意の文字列からMD5ハッシュ値を生成する処理例
	h := md5.New()
	io.WriteString(h, "ABCDE")

	//b := []byte{12, 34, 55, 3}

	fmt.Println(h.Sum(nil)) // [46 205 222 57 89 5 29 145 63 97 177 69 121 234 19 109]
	//fmt.Println(h.Sum(b))

	//fmt.Println(b)

	// %x : 16進数の文字列を生成
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s) // 2ecdde3959051d913f61b14579ea136d

}
