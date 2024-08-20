package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{2, 2, 0}, 4},
		{[]int{2, 21}, 23},
		{[]int{3, 21}, 24},
	}
	for _, v := range tests {
		x := sum(v.data...)
		if x != v.answer {
			t.Error("Expected ", v.answer, "Got", x)
		}
	}
}
