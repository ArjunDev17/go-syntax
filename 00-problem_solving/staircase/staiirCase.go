package main

import (
	"fmt"
	"strings"
)

// func printStair(count int) {
// 	for i := 0; i < count; i++ {
// 		for j := 0; j > i; j++ {
// 			if j > i {
// 				fmt.Print("#")
// 			} else {
// 				fmt.Print("-")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func staircase(n int32) {
	// Write your code here
	for i := int32(1); i <= n; i++ {
		// Cast int32 to int before passing to strings.Repeat
		fmt.Println(strings.Repeat(" ", int(n-i)) + strings.Repeat("#", int(i)))
	}
}
func main() {
	staircase(4)
}

//   *
//  **
// ***
