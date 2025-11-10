package main

import "fmt"

// Demonstrates different uses of the `for` loop in Go
func main() {
	standardForLoop()
	rangeForLoop()
	whileLikeLoop()
	doWhileLikeLoop()
}

// Standard for loop (like C-style)
func standardForLoop() {
	fmt.Println("Standard For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// Range-based loop (like foreach)
func rangeForLoop() {
	fmt.Println("Range For Loop:")
	arr := [5]int{1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}
}

// While-like loop using only the condition
func whileLikeLoop() {
	fmt.Println("While-Like Loop:")
	i := 0
	for i < 5 {
		fmt.Println("i:", i)
		i++
	}
}

// Do-while-like loop simulated using for + break
func doWhileLikeLoop() {
	fmt.Println("Do-While-Like Loop:")
	i := 0
	for {
		fmt.Println("i:", i)
		i++
		if i >= 5 {
			break
		}
	}
}

// first i wrote this code now i want to make it professional so i put my code in chatgpt
// package loop

// import "fmt"

// //for loop

// func main() {
// 	forLoopAsForLoop()
// 	forLoopAsForEachLoop()
// 	forLoopAsWhileLoop()
// 	forLoopAsdOWhileLoop()
// }
// func forLoopAsForLoop() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("%d :", i)
// 	}
// }
// func forLoopAsForEachLoop() {

// 	var arr [5]int

// 	arr = [5]int{1, 2, 3, 4, 5}

// 	for _, i := range arr {
// 		fmt.Println("values one by one :", i)
// 	}
// }
// func forLoopAsWhileLoop() {
// 	var i int
// 	for i < 10 {
// 		fmt.Println("i value incremetning :", i)
// 		i++
// 	}
// }
// func forLoopAsdOWhileLoop() {
// 	i := 4

// 	for {
// 		fmt.Println("it will com here")
// 		if i > 1 {
// 			fmt.Println("it will com here")
// 		}
// 	}()
// }
