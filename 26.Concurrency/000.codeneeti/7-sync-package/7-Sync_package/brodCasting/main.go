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
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) < 1 {
			// time.Sleep(1 * time.Millisecond)
			c.Wait()
		}
		fmt.Println("shard rsc 1 :", sharedRsc["rsc1"])
		c.L.Unlock()
	}()
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) < 2 {
			// time.Sleep(1 * time.Millisecond)
			c.Wait()
		}
		fmt.Println("shared resourc 2 :", sharedRsc["rsc2"])
		c.L.Unlock()
	}()
	c.L.Lock()
	//wrtie changes to shared rsc
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc1"] = "baar"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()

}
