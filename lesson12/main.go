package main

import (
	"fmt"
	f "fmt"
	"golangStudy/lesson12/foo"
)

/*
* go mod init : 依存モジュールを初期化
* go buildなどのビルドコマンド : 依存モジュールを自動インストール
* go list -m all : 現在の依存モジュールを表示
* go get : 依存モジュールの追加やバージョンアップ
* go mod tidy : 使われていない依存モジュールを削除
*
* 実は go mod を直接実行することは少なく、
* 他の go サブコマンドを実行したときに、自動的に処理が行われることが多い
*
* GO111MODULE = on
* goコマンドではモジュールを使用する必要があり、
* GOPATHを参照することはありません
* モジュール対応または「モジュール対応モード」で実行されるコマンドと呼びます。
*
* モジュール対応モードでは、GOPATHはビルド中のインポートの意味を定義しなくなりましたが、
* ダウンロードされた依存関係（GOPATH / pkg / mod）と
* インストールされたコマンド（GOBINが設定されていない場合はGOPATH / bin）
* を保存
*
* GO111MODULE = off
* goコマンドはモジュールサポートを使用しません
* ベンダーのディレクトリとGOPATHを調べて依存関係を見つけます
* 「GOPATHモード」と呼びます
*
* GO111MODULE = auto、または未設定
* goコマンドは現在のディレクトリに基づいてモジュールサポートを有効または無効にします
* モジュールのサポートは、現在のディレクトリにgo.modファイルが含まれる場合、
* またはgo.modファイルが含まれるディレクトリの下にある場合にのみ有効
 */

/*
--------------------------------------------------------------------------
* パブリックとプライベート、パッケージの分割（スコープ）
--------------------------------------------------------------------------
*/

// 関数のスコープ
func appName() string {
	const AppName = "GoApp"
	var Ver string = "1.0"
	return AppName + " " + Ver
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	// 関数内で重複して識別子を使いたい場合
	{
		b := "BBBB"
		fmt.Println(b) // BBBB
	}
	fmt.Println(b) // AAA
	return b       // AAA
}

func main() {
	// 一文字が大文字 = パブリック、外部のパッケージから参照可能
	f.Println(foo.Max) // 100
	// 一文字が小文字 = プライベート、外部のパッケージから参照不可
	// fmt.Println(foo.min) // undefined: foo.min
	f.Println(foo.ReturnMin()) // 1

	// 関数のスコープ
	fmt.Println(appName()) // GoApp 1.0
	// fmt.Println(AppName)   // undefined: AppName
	// fmt.Println(Ver)   		// undefined: Ver

	fmt.Println(Do("AAA"))
	// BBBB
	// AAA
	// AAA
}
