package main

import "fmt"

// UpdateAt updates the data of the node at a given position
func UpdateAt(head *Node, position int, newData interface{}) {
	if head.Next == nil {
		fmt.Println("List is empty.")
		return
	}

	current := head
	count := 0

	// Traverse to the node at the given position
	for current != nil && count < position {
		current = current.Next
		count++
	}

	// Check if the position is valid
	if current == nil {
		fmt.Println("Position out of range.")
		return
	}

	// Update the node's data
	current.Data = newData
}
