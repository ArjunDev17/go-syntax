package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	ex := make(chan int)

	go send(eve, odd, ex)
	recive(eve, odd, ex)
}
func recive(e, o, x <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel :", v)

		case v := <-o:
			fmt.Println("from the odd channel :", v)
		case v := <-x:
			fmt.Println("from x :", v)
			return
		}
	}
}

func send(e, o, ex chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	ex <- 0
}
