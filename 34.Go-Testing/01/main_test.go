package main

import "testing"

func TestSum(t *testing.T) {
	x := sum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, "Got", x)
	}
}
