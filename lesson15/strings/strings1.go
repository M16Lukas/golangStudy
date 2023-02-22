package main

import (
	"fmt"
	"strings"
)

/*
---------------------------------------------------------------
* strings
* 文字列操作機能がまとまったパッケージ
---------------------------------------------------------------
*/

func main() {
	// 文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2) // A,B,C ABC

	// 文字列に含まれる部分文字列を検索する
	// -1 = 含まれていない
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2) // 0 3

	// 部分文字列が複数含まれた場合、最後の部分文字列のインデックスが返される
	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3) // 5

	// 検索対象の文字列の中で第２引数の文字列のどれかが最初に現れるインデックスを返す
	i4 := strings.IndexAny("ABC", "ABC")
	// 検索対象の文字列の中で第２引数の文字列のどれかが最後に現れるインデックスを返す
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5) // 0 2

	// 検索対象の文字列が指定した文字列から開始されるかどうかをbool値で返す
	b1 := strings.HasPrefix("ABC", "A")
	// 検索対象の文字列が指定した文字列から終わるかどうかをbool値で返す
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2) // true true

	// 検索対象の文字列に指定した文字列が含まれているか
	b3 := strings.Contains("ABC", "B")
	// 検索対象の文字列の中で指定した文字列のいずれかの文字列が含まれているか
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4) // true true

	// 第１引数の文字列の中で第２引数の文字列が何回出現するか
	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "") // "" : 文字列の長さ + 1 を返す
	fmt.Println(i6, i7)               // 2 7

	// 文字列を繰り返して結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4) // ABCABCABCABC

	// 文字列の置換
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1) // -1 : 該当するすべての箇所を置換
	fmt.Println(s5, s6)                           // BAAAAA BBBBBB

	// 文字列を分割する
	// 返り値：[]string
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7) // [A B C D E]

	// SplitAfter : セパレーターを取り除かずに分割
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s8) // [A, B, C, D, E]

	// SplitN：分割する最大数を第３引数で指定
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s9) // [A B,C,D,E]

	// SplitAfterN : SplitAfterの最大数を指定
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s10) // [A, B, C, D,E]

	// 小文字に変換
	s11 := strings.ToLower("ABC")
	s12 := strings.ToLower("E")
	// 大文字に変換
	s13 := strings.ToUpper("abc")
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14) // abc e ABC E

	// 文字列から空白を取り除く
	s15 := strings.TrimSpace("    -    ABC    -    ")
	s16 := strings.TrimSpace("\tABC\r\n")
	// 全角も対象
	s17 := strings.TrimSpace("　　　　ABC　　　　")
	fmt.Println(s15, s16, s17) // -    ABC    - ABC ABC

	// 文字列からスペースで区切られたフィールドを取り出す
	// 返り値：[]string
	s18 := strings.Fields("a b c")
	fmt.Println(s18)    // [a b c]
	fmt.Println(s18[1]) // b
}
