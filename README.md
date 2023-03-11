# ゼロから始めるGolang勉強

## 教材
- [【Go入門】Golang基礎入門 + 各種ライブラリ + 簡単なTodoWebアプリケーション開発(Go言語)](https://www.udemy.com/course/golang-webgosql/)
- [現役シリコンバレーエンジニアが教えるGo入門 + 応用でビットコインのシストレFintechアプリの開発](https://www.udemy.com/course/go-fintech/)

## 内容

| フォルダ名 | タイトル |
| :--: | :--: |
| lesson1 | Goの基礎 |
| lesson2 | 変数 |
| lesson3 | 基本型 |
| lesson4 | 定数 |
| lesson5 | 演算子 |
| lesson6 | 関数 |
| lesson7 | 制御構文 |
| lesson8 | 参照型 |
| lesson9 | ポインタ |
| lesson10 | 構造体 |
| lesson11 | interface |
| lesson12 | パブリックとプライベートと分割 |
| lesson13 | テスト |
| lesson14 | Goのツールについて |
| lesson15 | 標準パッケージ |
| lesson16 | サードパーティーパッケージ |
| generics | Generics for Golang 1.18~ |
| parallel | 並列処理 |

## 各レッスンの詳細

#### 1. Goの基礎

#### 2. 変数
- 変数

#### 3. 基本型
- int(整数)型
- float(不動小数点)型
- uint、complex型
- bool(論理値)型
- string(文字列)型
- byte(uint8)型
- 配列型
- interface & nil 型
- 型変換

#### 4. 定数
- 定数

#### 5. 演算子
- 算術演算子
- 比較演算子
- 論理演算子
- シフト

#### 6. 関数
- 関数
- 無名関数
- 関数を返す関数
- 関数を引数に取る関数
- クロージャー
- ジェネレーター
- 可変長引数

#### 7. 制御構文
- if
- エラーハンドリング
- for
  - ラベル付きfor
- switch
  - switch型スイッチ
- defer
- panic & recover
- goroutin
- init

#### 8. 参照型
- スライス
  - append make len cap 
  - copy
  - for
  - 可変長引数
- map
  - for
- チャンネル(channel)
  - チャンネルとゴルーチン
  - close
  - for
  - select

#### 9. ポインタ
- ポインタ
- newとmakeの違い

#### 10. 構造体
- struct
  - メソッド
  - 埋め込み(Embedded) 
  - コンストラクタ
  - スライス
  - 独自型
- struct マップ

#### 11. interface
- interface
  - タイプアサーションとswitch type文
  - カスタムエラー
  - Stringer

#### 12. パブリックとプライベートと分割
- スコープ

#### 13. テスト
- テスト

#### 14. Goのツールについて
- Goコマンド

#### 15. 標準パッケージ
- os
- time
- math
- rand
- flag
- fmt
- log
- strconv
- strings
- bufio
- ioutil
- regexp
- sync
- crypto
- json
- sort
- context
- net
  - net/url
  - net/http client
  - net/http server

#### 16. サードパーティーパッケージ
- database/sqlite3
  - sqlite3
  - psql
- go-ini
- uuid

#### 17. Generics for Golang 1.18~
- Generics
  - interface
  - typesets
  - vector
  - struct
  - any
  - set

#### 18. 並列処理（更新中）
- Goroutine
- sync
  - WaitGroup
  - 競合 Mutex
  - デッドロック
  - リソース枯渇
  - MutexとRWMutex
  - Cond
  - ライブロック
  - Once
  - Pool
  - コネクションプールを使って、ベンチマークテスト
  - Map
  - Atomic
- Channel
  - 同期
  - close
  - buffer
  - ライフサイクルのカプセル化
  - チャネルの状態に対する操作の結果一覧
- Select
  - for-select
  - timeou
  - cancel
  - default

