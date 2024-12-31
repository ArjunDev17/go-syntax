package main

import "fmt"

func summitionOfWin(windowSize int, arr []int) []int {
	res := []int{}
	var windowSum int
	for i := 0; i < windowSize; i++ {
		windowSum += arr[i]
	}
	res = append(res, windowSum)

	for right := windowSize; right < len(arr); right++ {
		windowSum += arr[right]

		windowSum -= arr[right-windowSize]
		res = append(res, windowSum)

	}
	return res
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	windowSize := 4
	res := summitionOfWin(windowSize, arr)
	fmt.Println("res :", res)
}

// res : [6 9 12 15]
