package main

import "fmt"

/*
	TODO : Build a Pipeline
	genrate()->square()-->print
*/

// generator conver a list of integers to channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}

		close(out)
	}()
	return out
}

// squre - receive on inbound channel
// squre the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range in {
			out <- value * value
		}
		close(out)
	}()
	return out
}
func main() {
	//set up the pipeline

	//run the last stage of pipeline
	// recive the values from square stage
	// print each one,until channel is closed/.
	ch := generator(1, 2, 3, 4)
	out := square(ch)

	for res := range out {
		fmt.Println("squared data is :", res)
	}
	//composebility of pipeline
	for res := range square(square(generator(1, 2, 3))) {
		fmt.Println("squared data is 2:", res)
	}

}
