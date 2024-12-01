package main

// InsertFirstNode inserts a node at the beginning
func InsertFirstNode(data interface{}, head *Node) *Node {
	newNode := &Node{
		Data: data,
		Next: head.Next,
	}
	head.Next = newNode
	return head
}

// InsertLast inserts a node at the end
func InsertLast(data interface{}, head *Node) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}
	current := head
	// Traverse to the last node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func InsertAt(data interface{}, head *Node, position int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}
	count := 0 //1.  0<1
	prev := head
	for prev.Next != nil && count < position {
		prev = prev.Next
		count++
	}
	newNode.Next = prev.Next
	prev.Next = newNode

}
