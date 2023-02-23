package main

import (
	"bytes"
	"encoding/json"
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
* 応用編（POST）
* ヘッダーを付けたり、クエリを付けたりする場合
---------------------------------------
*/

func main() {
	// Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")
	// クエリの確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	// Postの時のデータ
	AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
	data, _ := json.Marshal(AccountData)

	// ResolveReference
	// 相対パスから絶対URLに変換する。
	// クエリをくっつけたURLを生成する。
	// baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint) // https://example.com/index/lists?id=1

	// bytes.NewBuffer([]byte("password")でリクエストの領域を作成
	// POSTの場合は、Bodyにデータを入れる。例えばパスワード。見られたらダメな情報はbodyに
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	// requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)
	// URLのクエリを確認
	q := req.URL.Query()
	// クエリを追加
	q.Add("name", "test")
	// クエリを表示
	fmt.Println(q) // map[id:[1] name:[test]]

	// &など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode()) // id=1&name=test
	// encodeしてからURLに戻す
	// 日本語などを変換する
	req.URL.RawQuery = q.Encode()

	// 実際にアクセスする
	// クライアントを作る
	var client *http.Client = &http.Client{}

	//結果 アクセスする
	resp, _ := client.Do(req)

	// 読み込み
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
