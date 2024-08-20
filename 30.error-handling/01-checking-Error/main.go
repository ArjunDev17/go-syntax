package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 6}
	lengthOfString, err := fmt.Print(arr) //Println(a ...any) (n int, err error)
	if err != nil {
		return
	}
	fmt.Println("\nlengthOfString :", lengthOfString)
}
