package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("value is :", <-ch)
}

// arjun@arjun-Vostro-3480:01_deadLockChallange$ go run main.go
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//         /media/arjun/863684ab-ea66-44f7-9b95-f624c9361dea1/GoLang/all-go/go-syntax/28.DeadLock/01_deadLockChallange/main.go:7 +0x37
// // exit status 2
