// Creates a slice of integers.
// create a go routine to calculate square of
// integer and store it in result slice and print it in fmt

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Print("Hello")
	numbers := []int{1, 2, 3, 4, 5}
	result := make([]int, len(numbers))

	var wg sync.WaitGroup

	for i, num := range numbers {
		wg.Add(1)

		go func(i, num int) {
			defer wg.Done()
			result[i] = num * num
		}(i, num)
	}

	wg.Wait()

	fmt.Println("Square : ", result)

}
