package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/*
-----------------------------------------
* net/http
* クライアント
-----------------------------------------
*/

func main() {

	//GETメソッド
	res, _ := http.Get("https://example.com")

	fmt.Println(res.StatusCode) // 200

	fmt.Println(res.Proto) // HTTP/2.0

	// ヘッダーの参照
	fmt.Println(res.Header["Date"])         // [Thu, 23 Feb 2023 06:37:57 GMT]
	fmt.Println(res.Header["Content-Type"]) // [text/html; charset=UTF-8]

	// リクエスト情報
	fmt.Println(res.Request.Method) // GET
	fmt.Println(res.Request.URL)    // https://example.com

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body) // レスポンスボディをすべて読み取る
	fmt.Print(string(body))             // ボディをstringに変換

	//------------------------------------
	// POSTメソッド
	// POSTに送信するデータを生成
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"

	res, err := http.PostForm("https://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))

}
