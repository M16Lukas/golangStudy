package alib

import "testing"

var Debug bool = false

// go test ./alib

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("スキップさせる")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Errorf("%v != %v", v, 3)
	}

}