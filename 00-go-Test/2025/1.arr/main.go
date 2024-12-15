package main

import "fmt"

func main() {
	arr := [5]int{2, 1, 6, 11, 3}

	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		} else if max < v {
			max = v
		}
	}
	fmt.Println("max:", max)
	fmt.Println("min :", min)
}
