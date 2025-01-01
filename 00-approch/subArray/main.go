package main

import "fmt"

func subArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			for k := i; k <= j; k++ {
				fmt.Print(arr[k], " ")
			}
			fmt.Println()
		}

	}
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	subArr(arr)
}

// i want to print submition of less than 10 and return the legth of the subarray
