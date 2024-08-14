/*
Synchronization:

Channels can be used to synchronize Go routines, ensuring that certain operations are completed before moving forward.
For example, a Go routine can wait to receive a signal from another Go routine before proceeding.
*/
package main

import (
	"fmt"
	"time"
)

func greet(name string, ch chan string) {
	fmt.Println("welcome from  :", name)
	time.Sleep(2 * time.Second) //simulate greeting time
	ch <- name + " greeted "    //send a msg to channel

}
func main() {
	guests := []string{"Ram", "Lakhan", "bharat", "Shatrughan"}
	ch := make(chan string)
	for _, guest := range guests {
		go greet(guest, ch)
	}
	for _, _ = range guests {
		msg := <-ch
		fmt.Println(msg)
	}
}
