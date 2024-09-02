package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()
	go func() {
		<-done
		<-done
		close(ch)
	}()
	for value := range ch {
		fmt.Println("data :", value)
	}
}

// go func() {
// 	<-done
// 	<-done
// 	close(ch)
// }()

// if above code we wrote like this it will not work check why

// <-done
// 	<-done
// 	close(ch)
