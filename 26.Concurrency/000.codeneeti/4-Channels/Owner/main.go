package main

import "fmt"

func main() {

	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}
	consumer := func(ch <-chan int) {
		//read value from channel
		for value := range ch {
			fmt.Println("value :", value)
		}
		fmt.Println("Done")
	}
	ch1 := owner()
	consumer(ch1)
}
