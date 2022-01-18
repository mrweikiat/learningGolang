package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{22, 20}, 42},
		test{[]int{1, 2, 3, 4, 5}, 15},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
