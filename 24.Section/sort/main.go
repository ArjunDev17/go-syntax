package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"rigved", "samved", "yajurved", "Atherved"}
	num := []int{3, 1, 33, 10, 9, 11}
	fmt.Println("sorted number :", num)
	sort.Ints(num)
	fmt.Println("sorted number :", num)
	fmt.Println("Unsorted String :", str)
	sort.Strings(str)
	fmt.Println("sorted String   :", str)

}
