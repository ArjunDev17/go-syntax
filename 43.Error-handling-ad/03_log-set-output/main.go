package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println(err)
		log.Println("error happened :", err)
	}
}
