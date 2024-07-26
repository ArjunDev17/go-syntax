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

type youngIn interface {
	walk()
	run()
}

func YoungRun(y youngIn) {
	y.run()
}

func main() {
	d1 := dog{
		first: "moti",
	}
	d1.walk()
	d1.run()
	// YoungRun(d1)

	d2 := &dog{
		first: "moti",
	}
	d2.walk()
	d2.run()
	YoungRun(d2)

}
