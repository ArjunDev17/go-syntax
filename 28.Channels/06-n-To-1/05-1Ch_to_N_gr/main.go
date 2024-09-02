package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 30; i++ {
			ch <- i
		}
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		go func() {
			for value := range ch {
				fmt.Println("data :", value)
			}
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

}
