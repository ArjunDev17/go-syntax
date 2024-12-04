package main

// // Add -sequential code to add number
// func Add(numbers []int64) int64 {
// 	var sum int64
// 	for _, number := range numbers {
// 		sum += int64(number)
// 	}
// 	return sum
// }

// how to make it faster
/**
Computing environment :
WHEN Add()is executed it runs on single core
Multi core processor ----[list of  cores in computer]but only one core is doing this task others are Idle

if we wamt to use Multi-Core Processoor
devide the inout and run multiple instances of add() function on each part parallel on diffrent core

*/

import (
	"fmt"
	"sync"
	"time"
)

// Add - Sequential code to add numbers
func Add(numbers []int64) int64 { //		{2,3,4,5,6,7}
	var sum int64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// AddParallel - Parallel code to add numbers
func AddParallel(numbers []int64) int64 {
	var sum int64
	var mu sync.Mutex // To safely update the sum variable
	var wg sync.WaitGroup

	// Divide the numbers slice into chunks for parallel processing
	numChunks := 4
	chunkSize := (len(numbers) + numChunks - 1) / numChunks

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}

		wg.Add(1)
		go func(nums []int64) {
			defer wg.Done()
			localSum := int64(0)
			for _, number := range nums {
				localSum += number
			}
			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait()
	return sum
}

func main() {
	// Example input
	numbers := make([]int64, 1_000_000)
	for i := range numbers {
		numbers[i] = int64(i + 1)
	}

	// Measure time taken for sequential addition
	start := time.Now()
	seqSum := Add(numbers)
	seqTime := time.Since(start)

	// Measure time taken for parallel addition
	start = time.Now()
	parSum := AddParallel(numbers)
	parTime := time.Since(start)

	// Output results
	fmt.Printf("Sequential Sum: %d, Time Taken: %v\n", seqSum, seqTime)
	fmt.Printf("Parallel Sum: %d, Time Taken: %v\n", parSum, parTime)
}
