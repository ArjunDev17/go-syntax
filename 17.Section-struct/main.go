package main

import "fmt"

type Name struct {
	first string
	last  string
}
type Student struct {
	Name       Name
	FatherName Name
}

func main() {
	// p1 := Name{
	// 	first: "Ram",
	// 	last:  "Ji",
	// }
	// // p2 := Name{
	// // 	first: "Sheeta",
	// // 	last:  "Ji",
	// // }
	// fmt.Printf("%T \t %#v\n", p1, p1)
	embededStruct()
	fmt.Println("----------------------------------")
	AnonymasStruct()
	composition()

}
func embededStruct() {
	p1 := Student{
		Name: Name{
			first: "Ram",
			last:  "Ji",
		},
		FatherName: Name{
			first: "Dev",
			last:  "rath",
		},
	}
	fmt.Printf("%+v :\n", p1)
}
func AnonymasStruct() {
	p1 := struct {
		first string
		last  string
	}{
		first: "Ram",
		last:  "Ji",
	}
	fmt.Printf("%+v :\n", p1)
}
func composition() {
	// example of illustrate composition in GoLang

}
