package main

import "fmt"

type Account interface {
	Balance() float64
}

type BankAccount struct {
	balance float64
}

// Value receiver - works for read-only methods
func (a BankAccount) Balance() float64 {
	return a.balance
}
func main() {
	var obj Account = BankAccount{balance: 45655}
	fmt.Println("res :", obj.Balance())
}
