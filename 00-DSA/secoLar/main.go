package main

import (
	"fmt"
	"math"
)

func secondLargest(arr []int) int {
	if len(arr) < 2 {
		fmt.Println("array must have greater then 2 elemtns ")
		return math.MinInt
	}
	laargest, second := math.MinInt, math.MinInt
	for _, num := range arr {
		if num > laargest {
			second = laargest
			laargest = num
		} else if num > second && num != laargest {
			second = num
		}

	}
	if second == math.MinInt {
		fmt.Println("no second largest element")
		return second
	}
	return second
}
func main() {
	arr := []int{10, 20, 4, 45, 99, 99}
	res := secondLargest(arr)
	if res != math.MinInt {
		fmt.Println("Second largest number :", res)
	}
}
