// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func (e *Employee) helperInit(name string, Age int, Salary float64) *Employee {
	e.Name = name
	e.Age = Age
	e.Salary = Salary
	return e
}
func main() {
	fmt.Println("Hello, 世界")
	var e Employee
	data := e.helperInit("arjun", 25, 75.)
	fmt.Printf("data :%+v\n", data)
	//helper function
}
