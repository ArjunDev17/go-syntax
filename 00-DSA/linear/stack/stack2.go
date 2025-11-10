package main

import "fmt"

var (
	stack = make([]int, 5)
)

func push1(stack1 *[]int) {

}
func stacks() {

	//declare an aray<which we will use as stack and will do implimentation for that>

	var choice int
	for {
		fmt.Println("enter your choice :")
		fmt.Println(" 1 push ")
		fmt.Println(" 2 pop")
		fmt.Println(" 3 peek")
		fmt.Println(" 4 display")
		fmt.Println("please enter your choice :")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			push1(&stack)
			fmt.Println("your data push")
		case 2:
			fmt.Println("your data pop")
		case 3:
			fmt.Println("your data peek")
		case 4:
			fmt.Println("your data display")

		}

	}

}
