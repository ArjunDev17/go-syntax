package main

import "testing"

type addTestCase struct {
	name     string
	a, b     int
	expected int
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []addTestCase{
		{name: "positive numbers", a: 5, b: 10, expected: 15},
		{name: "zero values", a: 0, b: 0, expected: 0},
		{name: "negative and positive", a: -3, b: 3, expected: 0},
		{name: "both negative", a: -2, b: -4, expected: -6},
		{name: "large numbers", a: 100000, b: 250000, expected: 350000},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
