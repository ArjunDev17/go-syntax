package main

import "fmt"

func main() {
	fmt.Println("for loop syntax")
	//basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i :", i)
	}
	//as while loop
	i := 0
	for i < 5 {
		fmt.Println("i :", i)
		i++
	}
	// // infinite loop
	// for {
	// 	fmt.Println("Long run")
	// }
	slice := []int{2, 4, 6, 8}
	for index, value := range slice {
		fmt.Printf("\n index :%d - value :%d", index, value)
	}
	maps := map[string]int{"Ram age": 108, "Lakhan age": 202}
	for key, value := range maps {
		fmt.Printf("\n key :%s - value :%d", key, value)

	}
}
