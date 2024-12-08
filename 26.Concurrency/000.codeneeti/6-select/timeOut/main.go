package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 108
	}()
	select {
	case m1 := <-ch1:
		fmt.Println("receve data in time :", m1)
	case <-time.After(2 * time.Second):
		fmt.Println("data did't recive because of late :")
	}
}

/*
arjun@arjun-Vostro-3480:timeOut$ go run main.go
receve data in time  108
arjun@arjun-Vostro-3480:timeOut$ go run main.go
data did't recive because of late :
arjun@arjun-Vostro-3480:timeOut$
*/
