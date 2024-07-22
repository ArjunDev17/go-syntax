package main

import "fmt"

type Customer struct {
	Name   string
	Salary int
}

func main() {
	var customer Customer
	fmt.Println("Customer Details :")
	fmt.Println("enter customer name :")
	fmt.Scanf("%s ", &customer.Name)
}
