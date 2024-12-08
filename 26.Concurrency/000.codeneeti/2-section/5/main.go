package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("value of i:", i)
// 		}()
// 	}
// 	wg.Wait()
// }

/*
value of i: 5
value of i: 5
value of i: 5
value of i: 5
value of i: 5
*/
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("value of i:", i)
		}(i)
	}
	wg.Wait()
}

/*
value of i: 4
value of i: 3
value of i: 1
value of i: 0
value of i: 2

*/
