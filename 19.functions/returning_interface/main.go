package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Barking"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "myauu"
}

func getAnimalVOice(animalType string) Animal {
	if animalType == "dog" {
		return Dog{}
	} else if animalType == "cat" {
		return Cat{}
	}
	return nil
}

func main() {

	fmt.Println("returning interface ")
	dog := getAnimalVOice("dog")
	fmt.Println("Dog voice :", dog.Speak())

	cat := getAnimalVOice("cat")
	fmt.Println("Cat noice :", cat.Speak())
}
