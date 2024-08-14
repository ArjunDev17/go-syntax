package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()

	}
	time.Sleep(1 * time.Second)
	fmt.Println("conunter vlaue is :", counter)
}

// go run -race main.go
// conunter vlaue is : 4702
// Found 3 data race(s)
// exit status 66
