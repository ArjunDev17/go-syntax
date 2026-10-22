package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printOdd(&wg)
	go printEven(&wg)

	wg.Wait()

	fmt.Println("Completed")
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even:", i)
		}
	}
}

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("Odd:", i)
		}
	}
}

func merge(intervals [][]int) [][]int {
	// Step 1: Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Step 2: Store merged intervals
	result := [][]int{}

	for _, interval := range intervals {
		n := len(result)

		// Step 3: Add interval if no overlap exists
		if n == 0 || result[n-1][1] < interval[0] {
			result = append(result, interval)
		} else {
			// Step 4: Merge overlapping intervals
			if interval[1] > result[n-1][1] {
				result[n-1][1] = interval[1]
			}
		}
	}

	return result
}
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return left
}