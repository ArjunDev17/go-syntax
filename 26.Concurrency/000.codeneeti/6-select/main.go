package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()
	//multiplex recv on channel
	for i := 0; i < 2; i++ {
		select {
		case c1 := <-ch1:
			fmt.Println("c1 realted to :", c1)
		case c2 := <-ch2:
			fmt.Println("C2 Related to :", c2)
		}
	}

}

/*

output diffrent because i changed sleep time

arjun@arjun-Vostro-3480:6-select$ go run main.go
c1 realte to : one
C2 Related to : two
arjun@arjun-Vostro-3480:6-select$ go run main.go
C2 Related to : two
c1 realted to : one
arjun@arjun-Vostro-3480:6-select$
*/
