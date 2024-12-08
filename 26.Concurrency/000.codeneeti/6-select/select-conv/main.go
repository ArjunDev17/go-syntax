package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 2; i++ {
			ch <- fmt.Sprintf("message : %d", i)
		}
		close(ch) // Close the channel to indicate no more messages will be sent.
	}()

	// Read from the channel without default to ensure data is processed.
	for {
		select {
		case m, ok := <-ch:
			if !ok {
				// Channel closed, exit loop.
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received:", m)
		default:
			fmt.Println("Waiting for data...")
			time.Sleep(500 * time.Millisecond) // Add delay to avoid busy-waiting.
		}
	}
}
