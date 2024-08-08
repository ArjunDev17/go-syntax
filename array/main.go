package main

import "fmt"

func main() {
	arr1 := []int{12, 1, 6, 7}
	// linearSearch(arr1, 4)
	// res := binerySearch(arr1, 4)
	// fmt.Println("res :", res)
	bubbleSort(arr1)
}
func linearSearch(arr []int, search int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			fmt.Printf("i :%d\n", arr[i])
		}
	}
}
func binerySearch(arr []int, search int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] > search {
			end = mid - 1
		} else if arr[mid] < search {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	for _, val := range arr {
		fmt.Printf("val :%d\n", val)
	}
}
