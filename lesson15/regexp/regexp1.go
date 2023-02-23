package main

import (
	"fmt"
	"regexp"
)

func main() {
	/* ------------------------------------------------------------
	* regexp
	* 正規表現による処理
	------------------------------------------------------------*/

	// regexp.MatchString　Goの正規表現の基本
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match) // true

	// Compile
	// 正規表現のパターンを予め生成しておく
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match) // true

	// MustCompile（推薦）
	// コンパイルエラーが発生した時際、直接ランタイムエラーを発生させる
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match) // true

	// 文字クラスも使用可能
	//regexp.MustCompile("\\d")
	//regexp.MustCompile(`\d`)

	// 正規表現のフラグ
	// フラグ一覧
	//
	// i 大文字小文字を区別しない
	// m マルチラインモード（＾と＆が文頭、文末に加えて行頭、行末にマッチ）
	// s .が\nにマッチ
	// U 最小マッチへの変換（x*はx*?へ、x+はx+?へ）
	re3 := regexp.MustCompile(`(?i)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match) // true

	// 幅を持たない正規表現のパターン
	// パターン一覧
	//
	// ^  文頭（mフラグが有効な場合は行頭にも）
	// $  文末（mフラグが有効な場合は行末にも）
	// \A 文頭
	// \z 文末
	// \b ASCIIによるワード協会
	// \B 非ASCIIによるワード協会
	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match) // true

	match = re4.MatchString("  ABC  ")
	fmt.Println(match) // false

	/*
		re := regexp.MustCompile("ab")
		re.MatchStrings("abc")
		//true
	*/

	re5 := regexp.MustCompile(".")
	match = re5.MatchString("ABC")
	fmt.Println(match) // true
	match = re5.MatchString("\n")
	fmt.Println(match) // false

	// 繰り返しを表す正規表現
	//
	// 繰り返しのパターン
	// x* ０回以上繰り返すｘ（最大マッチ）
	// x+ １回以上繰り返すｘ（最大マッチ）
	// x? ０回以上１回以下繰り返すｘ
	//
	// x{n,m} ｎ回以上ｍ回以下繰り返すｘ（最大マッチ）
	// x{n, } ｎ回以上繰り返すｘ（最大マッチ）
	// x{n}   ｎ回繰り返すｘ（最大マッチ）
	//
	// x*? ０回以上繰り返すｘ（最小マッチ）
	// x+? １回以上繰り返すｘ（最小マッチ）
	// x?? ０回以上１回以下繰り返すｘ（０回優先）
	//
	// x{n,m}? ｎ回以上ｍ回以下繰り返すｘ（最小マッチ）
	// x{n, }? ｎ回以上繰り返すｘ（最小マッチ）
	// x{n}?   ｎ回繰り返すｘ（最小マッチ）
	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))           // true
	fmt.Println(re6.MatchString("a"))            // true
	fmt.Println(re6.MatchString("aaaabbbbbbbb")) // true
	fmt.Println(re6.MatchString("b"))            // false

	re7 := regexp.MustCompile("A+?A+?X")
	fmt.Println(re7.MatchString("AAX"))        // ture
	fmt.Println(re7.MatchString("AAAAAAXXXX")) // true

	// 正規表現の文字クラス
	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y")) // true

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))     // true
	fmt.Println(re9.MatchString("abcdefg")) // false

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC")) // false
	fmt.Println(re10.MatchString("あ"))   // true

	// 正規表現のグループ
	// (正規表現)	グループ（順序によるキャプチャ）
	// (?:正規表現)	グループ（キャプチャされない）
	// (?:P<name>正規表現)	名前付きグループ
	re1 = regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re1.MatchString("abcxyz")) // true
	fmt.Println(re1.MatchString("ABCXYZ")) // true
	fmt.Println(re1.MatchString("ABCxyz")) // true
	fmt.Println(re1.MatchString("ABCabc")) // false

	// 正規表現にマッチした文字列の取得
	// FindString
	fs1 := re1.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1)                             // ABCXYZ
	fs2 := re1.FindAllString("ABCXYZABCXYZ", -1) // -1 : マッチしたすべての文字列を取得
	fmt.Println(fs2)                             // [ABCXYZ ABCXYZ]

	// 正規表現による文字列の分割
	// 正規表現にマッチした部分で文字列を分割する場合は、regexp.Regexp型のメソッドSplitを使う
	// 第二引数　分割する最大数を指定。-1でマッチした全ての箇所で分割する。[]stringで返す
	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1)) // [ASHVJV<H KNJBJVK ]
	re1 = regexp.MustCompile(`\s+`)
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1)) // [aaaaaaaaaa bbbbbbb cccccc]

	// 正規表現による文字列の置換
	// 正規表現にマッチした部分を別の文字列に置き換える
	// regexp.Regexp型メソッドReplaceAllString
	// 対象の文字列に正規表現のパターンにマッチする部分がない場合は、元の文字列がそのまま返される。
	// スペースを,に置き換える
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ",")) // AAA,BBB,CCC
	re1 = regexp.MustCompile(`、|。`)
	//また置換する文字列に空文字を指定することで、正規表現にマッチした部分を文字列から取り除ける
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", "")) // 私はGolangを使用するプログラマー

	// 正規表現のグループによるサブマッチ
	// FindAllStringSubmatch
	re1 = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	ms := re1.FindAllStringSubmatch(s, -1)
	for _, v := range ms {
		fmt.Println(v)
	}
	// [0123-456-7889 0123 456 7889]
	// [111-222-333 111 222 333]
	// [556-787-899 556 787 899]

	fmt.Println(ms[0][0]) // 0123-456-7889
	fmt.Println(ms[0][1]) // 0123
	fmt.Println(ms[0][2]) // 456

	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "あ-い-う"))

}
