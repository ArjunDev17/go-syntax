package main

import "fmt"

// PrintLL prints all nodes in the list
func PrintLL(head *Node) {
	current := head.Next
	if current == nil {
		fmt.Println("List is empty.")
		return
	}
	for current != nil {
		fmt.Println("Data:", current.Data)
		current = current.Next
	}
}
