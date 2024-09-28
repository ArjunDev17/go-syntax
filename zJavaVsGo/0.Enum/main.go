package main

import "fmt"

// Defining an enum using iota
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

func main() {
	var today Day = Wednesday
	fmt.Println("Today is:", today) // Output: Today is: Wednesday
}
