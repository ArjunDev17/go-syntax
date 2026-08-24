package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	num := 2
	ch := make(chan int)
	quit := make(chan os.Signal, 1)

	// Catch interrupt and terminate signals
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			num *= 2
			ch <- num
			time.Sleep(1 * time.Second)
		}
	}()

	// Listening for termination signal in a separate goroutine
	go func() {
		<-quit
		fmt.Println("\nProgram interrupted. Exiting gracefully...")
		os.Exit(0)
	}()

	for {
		fmt.Println(<-ch)
	}
}
