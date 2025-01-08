package main

type LinkedList struct {
	daat int
	Next *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	var prev *LinkedList
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
func main() {
	linkedList := &LinkedList{
		daat: 1,
	}
}
