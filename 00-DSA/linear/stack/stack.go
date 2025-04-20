package main

import "fmt"

func main() {
	// var stack []int
	stack := make([]int, 3)
	var choice int
loop:
	for {
		fmt.Println("enter your choice")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("please push your data insdie stack")
			var data int
			fmt.Println("enter your data")
			fmt.Scanf("%d", &data)
			push(&stack, data)
		case 2:
			fmt.Println("please chose pop options")
			pop(&stack)
		case 3:
			fmt.Println("please chose delete options")
			delete(&stack)
		case 4:
			fmt.Println("please chose display options")
			display(&stack)
		case 5:
			fmt.Println("Exiting the program...")
			break loop // breaks the outer for loop
		default:
		}
	}
}
func push(stack *[]int, data int) {
	*stack = append(*stack, data)
}
func pop(stack *[]int) {
	*stack = (*stack)[:len(*stack)-1]
}
func delete(stack *[]int) {

	*stack = (*stack)[:len(*stack)-1]
}
func display(stack *[]int) {
	fmt.Println("your whole data is :")
	for _, data := range *stack {
		fmt.Println(data)
	}
}
