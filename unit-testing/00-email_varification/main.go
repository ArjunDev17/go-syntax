package main

import (
	"fmt"
	"regexp"
)

func ValidateEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)

}
func main() {
	email := "arjun.devb25gmail.com"
	if ValidateEmail(email) {
		fmt.Println(email, "is a valid email")
	} else {
		fmt.Println(email, "is not valid")
	}
}
