package main

import "sync"

func main() {
	in := gen(2, 3)
	//Fan Out
	//Distribute the square work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)
	//Fan In
	//Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		println(n)
	}
}
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//func merge(a chan int, b chan int) chan int {
//	out := make(chan int)
//	go func() {
//		for {
//			select {
//			case n := <-a:
//				out <- n
//			case n := <-b:
//				out <- n
//			}
//		}
//	}()
//	return out
//
//}
