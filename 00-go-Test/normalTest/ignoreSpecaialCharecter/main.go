package main

import (
	"fmt"
	"unicode"
)

/*
input :i am going to meet my f{yriend
a

solution

	whenever if str[i]<=65 ||str[i]>=122){
			ch:=...
			println()
		}
		else{
			println("this charecter is not vzlid :",)
		}
*/
type userDefinedStr struct {
	Message string
	Gmail   string ` `
}

func ignoreSpecialChareter(cha string) {

}
func main() {
	var input rune

	fmt.Println("Please enter your text :")
	fmt.Scanf("%c", &input)

	// fmt.Printf("%s", str)

	if unicode.IsLetter(input) || unicode.IsDigit(input) {
		fmt.Println(input)
	} else {
		fmt.Println("this is a special chareter")
	}
}
