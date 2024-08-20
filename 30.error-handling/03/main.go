package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer f.Close()
	r := strings.NewReader("Ram is a GOD")
	io.Copy(f, r)
}
