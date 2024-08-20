package main

import (
	"fmt"
	"io"

	// "io/ioutil"
	"os"
)

func main() {
	po, err := os.Open("name.txt")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer po.Close()
	// bs, err := ioutil.ReadAll(po)
	bs, err := io.ReadAll(po)
	if err != nil {
		fmt.Println("err :", err)
	}
	fmt.Println("data :", string(bs))
}
