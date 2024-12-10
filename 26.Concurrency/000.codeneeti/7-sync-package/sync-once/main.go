package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup

// 	load := func() {
// 		fmt.Println("Run only once intialization func")
// 	}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			//Todo :modify so that load function gets called only once
// 			load()
// 		}()
// 	}
// 	wg.Wait()
// }
/*
arjun@arjun-Vostro-3480:sync-once$ go run main.go
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
Run only once intialization func
arjun@arjun-Vostro-3480:sync-once$
*/

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	load := func() {
		fmt.Println("Run only once intialization func")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//Todo :modify so that load function gets called only once
			// load()
			once.Do(load)
		}()
	}
	wg.Wait()
}

/*
arjun@arjun-Vostro-3480:sync-once$ go run main.go
Run only once intialization func
arjun@arjun-Vostro-3480:sync-once$
*/
