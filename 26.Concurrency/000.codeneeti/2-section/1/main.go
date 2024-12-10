package main

import (
	"fmt"
	"time"
)

func msg(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Call :", s)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	msg("Jai Shree Ram")
	/*
		TODO write goroutin with deffrent varients for function call .
		goroutine functioncall
		goroutine with anonymous function
		goroutine with function value call
		wait for goroutine to end

	*/

	// TODO write goroutine with deffrent varients for function call .

	// 	goroutine functioncall
	go msg("goroutine function call-2")

	// 	goroutine with anonymous function
	go func() {
		msg("goroutine Annonymous call-3 ")
	}()
	// 	goroutine with function value call

	fv := msg
	go fv("goroutine 4")
	// go func(str string) {
	// 	msg(str)
	// }("JaiShree")

	// 	wait for goroutine to end
	fmt.Println("waiting for goroutines.... ")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
