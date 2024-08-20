//arr 1,11,2,44,44,99,115,19,225 //

//output 115

package main

// func secondLargestnum(arr []int) int {
// 	if len(arr) {

// 	}
// 	max := arr[0]
// 	for index, _ := range arr {
// 		if max < arr[index] {
// 			max = arr[index]
// 		}
// 	}
// 	return max
// }

type Node struct {
	Data int
	Next *Node
}

func (node *Node) addNodeInLastMy(data int) {
	// hh
	head := &Node{}
	if head == nil {
		head.Data = data
		head.Next = head
	} else {
		for head != nil {
			head = head.Next
		}
		head.Data = data
		head.Next = node
	}
}

func (node *Node) addNodeInLast(data int) {
	newNode := &Node{Data: data}

	if node == nil {
		node = newNode
		return
	}

	current := node
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}
