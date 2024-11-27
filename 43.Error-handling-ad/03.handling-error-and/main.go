package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened :", err) //err happened : open no-file.txt: no such file or directory
		//log.Println("err happened", err) //2024/11/02 11:47:50 err happened open no-file.txt: no such file or directory
		//log.Fatal(err) //2024/11/02 11:49:44 open no-file.txt: no such file or directory
		//exit status 1
		panic(err) //panic: open no-file.txt: no such file or directory

	}
}
