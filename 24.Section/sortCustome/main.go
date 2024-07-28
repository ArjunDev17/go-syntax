package main

import (
	"fmt"
	"sort"
)

// Define the struct
type Person struct {
	Name string
	Age  int
}

// Define a type for sorting by Age
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// Define a type for sorting by Name
type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

// Define a type for sorting by Age and Name
type ByAgeAndName []Person

func (a ByAgeAndName) Len() int {
	return len(a)
}

func (a ByAgeAndName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAgeAndName) Less(i, j int) bool {
	if a[i].Age == a[j].Age {
		return a[i].Name < a[j].Name
	}
	return a[i].Age < a[j].Age
}

func main() {
	// Create a slice of Person
	people := []Person{
		{"Ram", 38},
		{"Lakhan", 25},
		{"Bharat", 22},
		{"Shatrughan", 20},
	}

	// Sort by age
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by Age:")
	for _, person := range people {
		fmt.Printf("%s: %d\n", person.Name, person.Age)
	}

	// Sort by name
	sort.Sort(ByName(people))
	fmt.Println("\nSorted by Name:")
	for _, person := range people {
		fmt.Printf("%s: %d\n", person.Name, person.Age)
	}

	// Sort by age and name
	sort.Sort(ByAgeAndName(people))
	fmt.Println("\nSorted by Age and Name:")
	for _, person := range people {
		fmt.Printf("%s: %d\n", person.Name, person.Age)
	}
}
