
package main

import "fmt"

func main() {

    fmt.Println("=== Example 1: continue statement ===")
    // continue -> skips the current iteration and moves to the next
    for i := 1; i <= 5; i++ {
        if i == 3 {
            // Skip number 3
            continue
        }
        fmt.Println("Value:", i)
    }

    fmt.Println("\n=== Example 2: break statement ===")
    // break -> exits the loop completely
    for i := 1; i <= 5; i++ {
        if i == 4 {
            // Stop the loop when i is 4
            break
        }
        fmt.Println("Value:", i)
    }

    fmt.Println("\n=== Example 3: goto statement ===")
    // goto -> jumps to a labeled section in code
    var num int = 1

    for num <= 5 {
        if num == 3 {
            goto skipLabel // Jump to label when num == 3
        }
        fmt.Println("Number:", num)
        num++
    }

skipLabel:
    fmt.Println("Jumped to skipLabel using goto")

    fmt.Println("\n=== Example 4: all three in one loop ===")
    // Combining break, continue, and goto for understanding
    for i := 1; i <= 10; i++ {
        if i == 3 {
            fmt.Println("Skipping 3 using continue")
            continue // skip this iteration
        }
        if i == 6 {
            fmt.Println("Jumping to label using goto")
            goto end // jump to end label
        }
        if i == 9 {
            fmt.Println("Breaking loop completely")
            break // exit the loop
        }
        fmt.Println("Processing value:", i)
    }

end:
    fmt.Println("Reached end label - program finished")
}
