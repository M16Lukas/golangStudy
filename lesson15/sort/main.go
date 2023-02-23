package main

import (
	"fmt"
	"sort"
)

/*
------------------------------------
* sort
------------------------------------
*/

// SliceとSliceStable
type Entry struct {
	Name  string
	Value int
}

func main() {

	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s) // [5 3 2 4 5 6 4 8 9 8 7 10] [a z j]

	// 並び順
	sort.Ints(i)    // asc(昇順)
	sort.Strings(s) // アルファベット順

	fmt.Println(i, s) // [2 3 4 4 5 5 6 7 8 8 9 10] [a j z]

	// SliceとSliceStable

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el) // [{A 20} {F 40} {i 30} {b 10} {t 15} {y 30} {c 30} {w 30}]

	// SliceとSliceStableの違い

	// Slice
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el) // [{b 10} {t 15} {A 20} {c 30} {i 30} {w 30} {y 30} {F 40}]
	fmt.Println("---------")

	// SliceStable
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("---------")
	fmt.Println(el) // [{b 10} {t 15} {A 20} {c 30} {i 30} {w 30} {y 30} {F 40}]
	fmt.Println("---------")
}
