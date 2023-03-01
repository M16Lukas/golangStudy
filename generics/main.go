package main

import (
	"fmt"
	"strconv"
)

// 従来の方法
func PrintSliceInts(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func PrintSliceStrings(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Generic：型を抽象化して、コードを再利用する
// any：あらゆる型を入れる、従来のinterface型のエイリアス型
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// interface : 型パラメータに制約を持たせる
type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func f[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

// typesets : 任意の型の集合を制約として持たせる
/*
 	~(チルダ) : 独自型（例：type MyInt int）の元となる型も満たしている時、
			インターフェースを満たしていると制約をかけたい場合
*/
type Number interface {
	~int | int32 | int64 | float32 | float64
}

func Max[T Number](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

// vector
type Vector[T any] []T
type IntVector = Vector[int]

// struct
type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

// メソッドは新たに型パラメータを持つことができない
// func (T[A, B, C]) m[U any](x U) {

// }

// set : 型パラメータを持つ型(Key, Value)
// comparable : 要素の比較が可能である制約
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, v := range xs {
		s.Add(v)
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {
	// 従来の方法
	PrintSliceInts([]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9})
	PrintSliceStrings([]string{"a", "as", "ae", "ag", "af", "s", "d", "b"})

	// Generic
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "as", "ae", "ag", "af", "s", "d", "b"})
	PrintSlice([]string{"a", "as", "ae", "ag", "af", "s", "d", "b"})

	// interface
	fmt.Println(f([]MyInt{1, 2, 3, 4})) // [1 2 3 4]

	// typesets
	fmt.Println(Max[int](1, 2))         // 2
	fmt.Println(Max[float64](1.1, 1.3)) // 1.3
	fmt.Println(Max(1.1, 1.3))          // 1.3

	// vector
	var v Vector[int] = Vector[int]{10, 20, 30}
	fmt.Println(v) // [10 20 30]

	var v2 Vector[float64] = Vector[float64]{1.3, 1.6, 5.6}
	fmt.Println(v2) // [1.3 1.6 5.6]

	v3 := IntVector{1, 2, 3}
	fmt.Println(v3) // [1 2 3]

	// struct
	var t T[int, []*int, *int]
	fmt.Printf("A: %T, B: %T, C: %T\n", t.a, t.b, t.c) // A: int, B: []*int, C: *int

	// any : 従来のinterface型のエイリアス型
	var a any = 1
	a = "a"
	a = true
	fmt.Println(a) // true

	// set : 型パラメータを持つ型(Key, Value)
	s := NewSet(1, 2, 3)
	fmt.Println(s) // map[1:{} 2:{} 3:{}]
	s.Remove(1)
	fmt.Println(s) // map[2:{} 3:{}]
}
