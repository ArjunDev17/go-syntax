package main

import "fmt"

type Animal interface {
	Voice()
}
type Dog struct{}
type Cat struct{}

func (d Dog) Voice() {
	fmt.Println("voice type")
}
