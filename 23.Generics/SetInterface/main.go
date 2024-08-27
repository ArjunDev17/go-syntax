package main

import (
	"fmt"
	"strings"
)

type myNum interface {
	int | float64
}

func genericsCal[T myNum](a, b T) T {
	return a - b
}
func main() {

	fmt.Println("--------------------")
	fmt.Println("by generics :", genericsCal(4, 2))
	fmt.Println("by generics :", genericsCal(4.2, 2.1))
}
func reverseWords(s string) string {
	words := strings.Fields(s)
	i := 0
	j := len(words) - 1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}
