/*
Design and implement an In Memory Loan EMI Calculator.

The code should have functionality to create users. Users can be either customer or admin. All users will have a unique identifier: username.
Admin can create a Loan in the system for a customer.
While creating a loan, admin_username, customer_username, principal amount, interest rate and time (loan tenure) need to be taken as input.
The interest for the loan is calculated by I = (P*N*R)/100 where P is the principal amount, N is the number of years and R is the rate of interest. The total amount to repay will be A = P + I The amount should be paid back monthly in the form of EMIs. Each EMI = A/(N*12)
Users should be able to make EMI payments for their loan only.
Users should be able to fetch loan info for their loans only. Fetching a loan should return the loan info along with all the emi payments done and EMIs remaining[optional].
Admin should be able to fetch all loans for all customers.
All the functions should take username as one of the arguments, and all user level validation should happen against this username.
*/
package main

import (
	"errors"
	"fmt"
	"math"
)

type User struct {
	Username string
	Isadmin  bool
}

type Loan struct {
	admin_username    string
	customer_username string
	principal         float64
	interest_rate     float32
	time              int
	emi               float64
	totalAmount       float64
}

type InmemoryStore struct {
	Users map[string]User
	Loans map[string][]Loan
}

func main() {
	store := InmemoryStore{
		Users: make(map[string]User),
		Loans: make(map[string][]Loan),
	}

	store.CreateUser("admin", true)
	store.CreateUser("Arjun", false)

	value := store.CreateLoan("Arjun", 10000, 7.5, 5)
	fmt.Println("err :", value)
	fmt.Println("stored Loan :", store)

}

func (store *InmemoryStore) CreateUser(username string, isAdmin bool) error {
	if _, exists := store.Users[username]; exists {
		return errors.New("User already exists")
	}
	store.Users[username] = User{Username: username, Isadmin: isAdmin}
	return nil
}

func (store *InmemoryStore) CreateLoan(username string, principal, rate float64, time int) Loan {
	ratePerMonth := rate / (12 * 100)
	month := time * 12
	emi := (principal * ratePerMonth * math.Pow(1+ratePerMonth, float64(month))) / (math.Pow(1+ratePerMonth, float64(float64(month))) - 1)
	totalAmount := emi * float64(month)
	loan := Loan{
		admin_username: username,
		principal:      principal,
		interest_rate:  rate,
		time:           time,
		emi:            emi,
		totalAmount:    totalAmount,
	}
	store.Loans[username] = append(store.Loans[username], loan)
	return loan
}
