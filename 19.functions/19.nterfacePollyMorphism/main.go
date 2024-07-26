package main

import "fmt"

type person struct {
	first string
}
type agent struct {
	person
	age int
}

func (p person) speak() {
	fmt.Println(p.first, "speek like his age :")
}
func (a agent) speak() {
	fmt.Println(a.first, "speek like his age :", a.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	fmt.Println("")
	agent1 := agent{
		person: person{
			first: "Ram",
		},
		age: 108,
	}
	agent2 := agent{
		person: person{
			first: "Ram Ji",
		},
		age: 108,
	}
	fmt.Println("1 way to call---------------")
	agent1.speak()
	agent2.speak()

	fmt.Println("2 way to call---------------")
	saySomething(agent1)
	saySomething(agent2)
}
