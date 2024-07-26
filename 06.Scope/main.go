// package main

// import "fmt"

// var x = 12 // package-level variable

// func main() {
// 	x := 11 // function-level variable
// 	fmt.Printf("\nfunction-level x: %d", x)

// 	{
// 		fmt.Printf("\nfunction-level x from inner block: %d", x) // Access function-level variable directly

// 		// Access package-level variable by referring to the global scope
// 		fmt.Printf("\npackage-level x: %d", main.x)

//			// Declare a block-level variable after accessing the outer variables
//			x := 10 // block-level variable
//			fmt.Printf("\nblock-level x: %d\n", x)
//		}
//	}
package main

import "fmt"

var x = 12 // package-level variable

func main() {
	x := 11 // function-level variable
	fmt.Printf("\nfunction-level x: %d", x)

	{
		fmt.Printf("\nfunction-level x from inner block: %d", x) // Access function-level variable directly

		// Access package-level variable using a new block to avoid shadowing
		{
			fmt.Printf("\npackage-level x: %d", x)
		}

		// Declare a block-level variable after accessing the outer variables
		x := 10 // block-level variable
		fmt.Printf("\nblock-level x: %d\n", x)
	}
	if z := 20; z > 10 {
		fmt.Println("z value :", z)
	}
}
