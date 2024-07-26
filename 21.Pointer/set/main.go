package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println(d.first, "good name")
}
func (d *dog) run() {
	fmt.Println(d.first, ":like milk")
}

func main() {
	d1 := dog{
		first: "moti",
	}
	d1.walk()
	d1.run()

	d2 := &dog{
		first: "moti",
	}
	d2.walk()
	d2.run()
}
