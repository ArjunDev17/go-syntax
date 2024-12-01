package main

import "fmt"

// DeleteAt deletes a node at a specific position
func DeleteAt(head *Node, position int) *Node {
	if head.Next == nil {
		fmt.Println("List is empty. Cannot delete.")
		return head
	}

	current := head
	count := 0

	// Traverse to the node just before the one to be deleted
	for current != nil && count < position-1 {
		current = current.Next
		count++
	}

	// Check if the position is valid
	if current == nil || current.Next == nil {
		fmt.Println("Position out of range.")
		return head
	}

	// Skip the node at position
	current.Next = current.Next.Next
	return head
}

// DeleteFirst deletes the first node
func DeleteFirst(head *Node) *Node {
	if head.Next == nil {
		fmt.Println("List is empty. Cannot delete.")
		return head
	}
	// Move the head pointer to the second node
	head.Next = head.Next.Next
	return head
}

// DeleteLast deletes the last node
func DeleteLast(head *Node) {
	if head.Next == nil {
		fmt.Println("List is empty. Cannot delete.")
		return
	}
	current := head
	// Traverse to the second-to-last node
	for current.Next != nil && current.Next.Next != nil {
		current = current.Next
	}
	// Remove the last node
	current.Next = nil
}
