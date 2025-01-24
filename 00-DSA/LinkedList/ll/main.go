package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

func CreatNode(data int, node *Node) Node {
	return Node{
		data: data,
		next: node,
	}
}
func main() {
	var head Node
	for {
		var choice int
		fmt.Println("please enter your choice here ")
		fmt.Scanf("%d", &choice)
		if choice == 3 {
			os.Exit(1)
		}
	}
	head = CreatNode(10, &head)
	fmt.Printf("\n%+v\n", head)
	fmt.Printf("\n%+v\n", head.next)

	head2 := CreatNode(20, &head)
	head.next = &head2

	fmt.Printf("\n%+v\n", head2)
	fmt.Printf("\n%+v\n", head2.next)
	head3 := CreatNode(30, &head2)
	head.next = &head3

	fmt.Printf("\n%+v\n", head3)
	fmt.Printf("\n%+v\n", head3.next)

}
