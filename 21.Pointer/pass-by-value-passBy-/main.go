package main

import "fmt"

func inDelta(n *int) {
	*n = 23
}
func arDelta(a []int) {
	a[0] = 23
}
func main() {
	a := 12
	fmt.Println("a :", a)
	inDelta(&a)
	fmt.Println("a :", a)
	fmt.Println("--------------------")
	ar := []int{1, 2, 3, 4}
	fmt.Println("ar :", ar)
	arDelta(ar)
	fmt.Println("ar :", ar)
	fmt.Println("--------------------")
	m := make(map[string]int)
	m["name"] = 34
	fmt.Println("m: ", m)
	mapDelta(m, "name")
	fmt.Println("m: ", m)
}
func mapDelta(m1 map[string]int, s string) {
	m1[s] = 44
	fmt.Println("m1 :", m1)
}
