package main

import (
	// "github.com/crate-crypto/go-ipa/bandersnatch/fp"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	str := "arjun@123"
	bp, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		fmt.Println("you can't login", err)
	}
	fmt.Println("your password :", bp)
	pwd := "arjun@123"
	err = bcrypt.CompareHashAndPassword(bp, []byte(pwd))
	if err != nil {
		fmt.Println("you can't login")
		return
	}
	fmt.Println("you are login")

}
