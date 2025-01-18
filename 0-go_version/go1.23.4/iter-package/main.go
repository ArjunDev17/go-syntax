package main

import "fmt"

// Define a custom iterator for a range of integers.
type IntRange struct {
	current, end int
}

// Next retrieves the next value in the range.
func (r *IntRange) Next() (int, bool) {
	if r.current < r.end {
		value := r.current
		r.current++
		return value, true
	}
	return 0, false // No more elements.
}

// NewIntRange initializes the range iterator.
func NewIntRange(start, end int) *IntRange {
	return &IntRange{current: start, end: end}
}

func main() {
	// Create an IntRange iterator from 1 to 5.
	intRange := NewIntRange(1, 5)

	// Use the custom iterator.
	for {
		value, ok := intRange.Next()
		if !ok {
			break
		}
		fmt.Println(value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
