package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	name string
}

func (b Book) String() string {
	return fmt.Sprint(" :", b.name)
}

type count int

func (c count) String() string {
	return fmt.Sprint(" this is the number-:", strconv.Itoa(int(c)))
}
func logInfo(s fmt.Stringer) {
	log.Print("Log from local/ :", s.String())
}
func main() {
	b := Book{
		name: "The Ramayan",
	}
	var c count = 12
	fmt.Print(b)
	fmt.Print(c)
	fmt.Println("------------------------------")
	log.Print(b)
	log.Print(c)
	fmt.Println("------------------------------")
	logInfo(b)
	logInfo(c)
}
