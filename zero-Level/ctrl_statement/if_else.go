package main

import (
	"fmt"
)

func ifmain() {
	fmt.Println("today we are going to see if else statement")
	var number int
	fmt.Println("enter your number :")
	fmt.Scanf("%d ", &number)
	if number > 12 {
		fmt.Println("number is greater than 12")
	} else {
		fmt.Println("number is less than 12")
	}
}
