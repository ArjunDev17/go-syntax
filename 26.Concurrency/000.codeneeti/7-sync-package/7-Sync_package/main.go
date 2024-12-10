package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()
		//TODO :suspend goroutine until sharedRsc is populated
		c.L.Lock()
		for len(sharedRsc) == 0 {
			// time.Sleep(1 * time.Millisecond)
			c.Wait()
		}
		fmt.Println("Shared Resources :", sharedRsc["rsc1"])
		c.L.Unlock()
	}()
	c.L.Lock()
	//write changes to shared resources
	sharedRsc["rsc1"] = "foo"
	c.L.Unlock()
	wg.Wait()
	fmt.Println("done.....")

}
