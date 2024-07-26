package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Sub  []string `json:"sub"`
}

func main() {
	stuData := Student{
		Name: "Ram",
		Age:  108,
		Sub:  []string{"yajurved", "samved", "atherved", "rigved"},
	}
	b, err := json.Marshal(stuData)
	if err != nil {
		log.Fatal(err)
	}
	// Print the compact JSON
	os.Stdout.Write(b)
	fmt.Println("\n--------------------")
	// Marshal the struct to pretty JSON
	b1, err := json.MarshalIndent(stuData, "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Print the pretty JSON
	os.Stdout.Write(b1)
	useOfUnMarshal()
}

func useOfUnMarshal() {
	fmt.Println("\n--------------------------------------")
	var stu Student
	jsonStr := `{
		"name": "Ram",
		"age": 108,
		"sub": [
			"yajurved",
			"samved",
			"atherved",
			"rigved"
		]
	}`
	err := json.Unmarshal([]byte(jsonStr), &stu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", stu)
}
