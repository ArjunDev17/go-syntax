// function
package main

import "fmt"

func main() {
	str := "tesat1"
	str1 := "test 2"
	go concateStr(str, str1)

	// a-b
	// rateLimiter
	// fmt.Println("concate string are :", res)
}
func concateStr(s1, s2 string) {
	s3 := s1 + s2
	fmt.Println("s3 :", s3)
}
