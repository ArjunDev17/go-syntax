package main

import "fmt"

type Dog struct{}
type Cat struct{}

func (d Dog) say() string {
	return "Bark"
}
func (c Cat) say() string {
	return "myau"
}
func main() {
	c := Cat{}
	d := Dog{}
	fmt.Printf("Cat %s \n", c.say())

	fmt.Printf("Dog noice :%s\n", d.say())

}
