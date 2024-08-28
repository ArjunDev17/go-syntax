package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Ram"] = 108
	ages["Bharat"] = 107
	ages["Shatrughan"] = 106
	ages["Lakshman"] = 105

	fmt.Println("Ram Ages is :", ages["Ram"])
	age, exist := ages["Ravan"]
	if exist {
		fmt.Println("Ravan ages is :", age)
	} else {
		fmt.Println("Ravan age is not exist")
	}
}
