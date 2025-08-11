package main

import "fmt"

type Account interface {
	Balance() float64
	Deposite(float64) error
}
type MyBank struct {
	balance float64
}

func (ac *MyBank) Balance(amount float64) float64 {
	return ac.balance
}
func (ac *MyBank) Deposite(amount float64) error {
	if amount < 0 {
		fmt.Println("amount is less than 0")
	}
	ac.balance += amount
	return nil
}
func main() {
	account := &MyBank{}
	account.Deposite(50)
	err := account.Deposite(500)
	if err != nil {
		fmt.Println("Error !", err)
	} else {
		fmt.Println("Current amount is :", account.balance)
	}

}
