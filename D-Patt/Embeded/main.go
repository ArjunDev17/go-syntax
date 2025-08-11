package main

import "fmt"

type Account interface {
	Balance(float64) float64
	Deposite(float64) error
}
type BasicAccount struct {
	balance float64
}

func (ac *BasicAccount) Balance(amount float64) float64 {
	return ac.balance
}
func (ac *BasicAccount) Deposite(amount float64) error {
	if amount < 0 {
		fmt.Println("amount is less than 0")
	}
	ac.balance += amount
	return nil
}
