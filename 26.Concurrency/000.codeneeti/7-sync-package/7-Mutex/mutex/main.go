package main

import (
	"fmt"
	"sync"
)

func main() {
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance += amount
	}

	withdraw := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	// Correctly manage wait groups
	wg.Add(200)

	// Deposits
	for i := 0; i < 100; i++ {
		go func(amount int) {
			defer wg.Done()
			deposit(amount)
		}(i) // Pass `i` as an argument to capture its value
	}

	// Withdrawals
	for i := 0; i < 100; i++ {
		go func(amount int) {
			defer wg.Done()
			withdraw(amount)
		}(i) // Pass `i` as an argument to capture its value
	}

	wg.Wait()
	fmt.Printf("Final Balance: %d\n", balance)
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var balance int

// 	var wg sync.WaitGroup

// 	var mu sync.Mutex

// 	deposit := func(amount int) {
// 		mu.Lock()
// 		balance += amount
// 		mu.Unlock()
// 	}
// 	withdrawl := func(amount int) {
// 		mu.Lock()
// 		defer mu.Unlock()
// 		balance -= amount
// 	}
// 	go func() {

// 	}()

// 	wg.Add(100)

// 	//Deposite
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			defer wg.Done()
// 			deposit(i)
// 		}()

// 	}
// 	//WithDrwal
// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			defer wg.Done()
// 			withdrawl(i)
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Printf("wg amt: %d\n", balance)
// }
