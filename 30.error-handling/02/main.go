package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string
	fmt.Println("your name :")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}
	fmt.Println("your role  :")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}
	fmt.Println("your lanuage :")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}
	fmt.Println("your name :", ans1, ans2, ans3)
}
