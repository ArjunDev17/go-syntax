package loop

import "fmt"

//for loop

func main() {
	forLoopAsForLoop()
	forLoopAsForEachLoop()
	forLoopAsWhileLoop()
	forLoopAsdOWhileLoop()
}
func forLoopAsForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d :", i)
	}
}
func forLoopAsForEachLoop() {

	var arr [5]int

	arr = [5]int{1, 2, 3, 4, 5}

	for _, i := range arr {
		fmt.Println("values one by one :", i)
	}
}
func forLoopAsWhileLoop() {
	var i int
	for i < 10 {
		fmt.Println("i value incremetning :", i)
		i++
	}
}
func forLoopAsdOWhileLoop() {
	i := 4

	for {
		fmt.Println("it will com here")
		if i > 1 {
			fmt.Println("it will com here")
		}
	}
}
