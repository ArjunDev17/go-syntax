package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func main() {
	head := &Node{}

	// Menu for user interaction
	var choice int
	for {
		fmt.Println("\nLinked List Operations:")
		fmt.Println("1: Insert at First Position")
		fmt.Println("2: Insert at Last Position")
		fmt.Println("3: Display List")
		fmt.Println("4: Delete First Node")
		fmt.Println("5: Delete Last Node")
		fmt.Println("6: Delete at Specific Position")
		fmt.Println("7: Update at Specific Position")
		fmt.Println("8: Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var data string
			fmt.Print("Enter data to insert at the first position: ")
			fmt.Scan(&data)
			head = InsertFirstNode(data, head)

		case 2:
			var data string
			fmt.Print("Enter data to insert at the last position: ")
			fmt.Scan(&data)
			InsertLast(data, head)

		case 3:
			fmt.Println("Displaying List:")
			PrintLL(head)

		case 4:
			head = DeleteFirst(head)

		case 5:
			DeleteLast(head)

		case 6:
			var position int
			fmt.Print("Enter the position to delete: ")
			fmt.Scan(&position)
			head = DeleteAt(head, position)

		case 7:
			var position int
			var newData string
			fmt.Print("Enter position to update: ")
			fmt.Scan(&position)
			fmt.Print("Enter new data: ")
			fmt.Scan(&newData)
			UpdateAt(head, position, newData)

		case 8:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
