package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // Buffered channel with size 3

	// Producer goroutine
	go func() {
		for i := 1; i <= 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i // Send value into the channel
			fmt.Printf("Sent: %d\n", i)
		}
		close(ch) // Close the channel after sending all values
	}()

	// Delay the consumer to allow buffer to fill up
	time.Sleep(2 * time.Second)

	// Consumer
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
		time.Sleep(1 * time.Second) // Simulate slow consumption
	}
}
