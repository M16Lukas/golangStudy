package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}
type List []Entry

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j
	Swap(i, j int)
}

// ---
func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここをカスタムする
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

//---

func main() {

	m := map[string]int{"ada": 1, "hoge": 4, "basha": 3, "poeni": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	//sort.SliceStable(lt, func(i, j int) bool { return lt[i].Name < lt[j].Name })
	//fmt.Println(lt)

	// sort
	sort.Sort(lt)
	fmt.Println(lt) // [{ada 1} {basha 3} {poeni 3} {hoge 4}]

	// Reverse
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt) // [{hoge 4} {poeni 3} {basha 3} {ada 1}]

}
