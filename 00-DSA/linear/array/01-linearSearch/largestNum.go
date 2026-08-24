package main

import "fmt"

func largestNum(num []int) {
	minN, maxN := 0, 0
	for i := range num {
		if minN < num[i] {
			maxN = num[i]
		}
	}
	fmt.Println("maximum number :", maxN)
}
