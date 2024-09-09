package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	// for {
	// 	fmt.Println("data is :", <-ch)
	// }
	for value := range ch {
		fmt.Println("data is :", value)
	}
}
