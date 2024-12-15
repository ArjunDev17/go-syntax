package main

import (
	"fmt"
	"sync"
)

func calCulate(a, b int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for i := a; i < b; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	resultChan <- sum

}

func main() {

	resultChan := make(chan int, 3)
	var wg sync.WaitGroup

	eveRange := [3][2]int{
		{1, 33},
		{34, 64},
		{64, 100},
	}

	for _, v := range eveRange {
		wg.Add(1)
		go calCulate(v[0], v[1], resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	totalSum := 0
	for c := range resultChan {
		totalSum += c
	}
	fmt.Println("total Sum :", totalSum)
}
