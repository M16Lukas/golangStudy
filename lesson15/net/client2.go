package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Account struct {
	ID       string
	PassWord string
}

/*
---------------------------------------
* net/http clientの応用
* 応用編（GET）
* ヘッダーを付けたり、クエリを付けたりする場合
---------------------------------------
*/

//get

func main() {
	// 応用
	// ヘッダーをつけたり、クエリをつけたり
	// Parse　正しいURLか確認
	base, _ := url.Parse("https://example.com/")

	// クエリ　の確認。URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	// ResolveReference
	// クエリをくっつけたURLを生成する。
	// 相対パスから絶対URLに変換する。
	// baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint) // https://example.com/index/lists?id=1

	// GET ver
	// リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	// まだリクエストはしていない。structを作っただけ。
	req, _ := http.NewRequest("GET", endpoint, nil)

	// requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)

	// URLのクエリを確認
	q := req.URL.Query()
	// クエリを追加
	q.Add("name", "test")
	// クエリを表示
	fmt.Println(q) // map[id:[1] name:[test]]
	// &など特殊文字などがある場合があるため、
	// encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode()) // id=1&name=test

	// エンコードしてからURLに戻す
	// 日本語などを変換する
	req.URL.RawQuery = q.Encode()

	// 実際にアクセスする
	// クライアントを作る
	var client *http.Client = &http.Client{}

	// 結果、アクセスする
	resp, _ := client.Do(req)

	// 読み取り
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
