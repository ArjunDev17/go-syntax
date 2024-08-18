// package main

// import (
// 	"fmt"
// )

// type LinkedList struct {
// 	Number int
// 	Next   *LinkedList
// }

// // create node
// func (node *LinkedList) createNode(num int) {
// 	var temp = &LinkedList{}
// 	temp.Number = num
// 	temp.Next = nil
// 	if node == nil {
// 		node = temp
// 	} else {
// 		p := &LinkedList{}
// 		for p.Next == nil {
// 			p = p.Next
// 		}
// 		p = temp
// 	}
// }
// func (node *LinkedList) displayNode() {
// 	p := node
// 	for p != nil {
// 		fmt.Printf("Data is ->%d", p.Number)
// 	}
// }
// func main() {
// 	// head:=&LinkedList{}
// 	head := new(LinkedList)
// 	head.createNode(10)
// 	head.createNode(10)
// 	head.createNode(10)
// 	head.createNode(10)
// 	head.displayNode()

// }
package main

import (
	"fmt"
)

type LinkedList struct {
	Number int
	Next   *LinkedList
}

// createNode adds a new node with the given number at the end of the list
func (node *LinkedList) createNode(num int) {
	temp := &LinkedList{Number: num, Next: nil}

	if node == nil {
		node = temp
	} else {
		p := node
		for p.Next != nil {
			p = p.Next
		}
		p.Next = temp
	}
}

// displayNode prints the data in each node of the list
func (node *LinkedList) displayNode() {
	p := node
	for p != nil {
		fmt.Printf("Data is -> %d\n", p.Number)
		p = p.Next
	}
}

func main() {
	head := &LinkedList{Number: 10}
	head.createNode(20)
	head.createNode(30)
	head.createNode(40)
	head.displayNode()
}
