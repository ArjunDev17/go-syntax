package main

import (
	"fmt"
)

// user Defined
func forSlipt(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			count++
		}
	}
	return count
}
func countParagraph(str string) int {
	n := forSlipt(str)
	// n := len(strings.Split(str, ".")) - 1
	return n
}
func main() {
	str := "abcs.bcs.lklkjldsa.jhk"
	count := countParagraph(str)
	fmt.Println("count :", count)
}
