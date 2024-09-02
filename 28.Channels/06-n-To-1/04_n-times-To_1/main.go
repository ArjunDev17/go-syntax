package main

import "fmt"

func main() {
	count := 10
	ch := make(chan int)
	done := make(chan bool)

	for i := 0; i < count; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
			done <- true
		}()
	}
	go func() {
		for i := 0; i < count; i++ {
			<-done
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Println("data :", value)
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	count := 10
// 	ch := make(chan int)
// 	done := make(chan bool)
// 	semaphore := make(chan struct{}, count) // Semaphore with buffer size equal to the number of goroutines

// 	for i := 0; i < count; i++ {
// 		semaphore <- struct{}{} // Acquire a spot in the semaphore
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				ch <- i
// 			}
// 			<-semaphore // Release the spot in the semaphore
// 			done <- true
// 		}()
// 	}

// 	// Close the channel once all goroutines are done
// 	go func() {
// 		for i := 0; i < count; i++ {
// 			<-done
// 		}
// 		close(ch)
// 	}()

// 	// Read from the channel
// 	for value := range ch {
// 		fmt.Println("data:", value)
// 	}
// }
