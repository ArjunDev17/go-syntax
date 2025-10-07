package ifelse

import (
	"fmt"
)

func main() {
	useOfIf_else()
	useOfIf_elseIf()
	useIfandShortHand()
}
func useOfIf_else() {
	fmt.Println("use of if else")
	num1 := 10
	if num1 > 12 {
		fmt.Println("num1 greater")
	} else {
		fmt.Println("num1 less")
	}

}
func useOfIf_elseIf() {
	a, b, c := 12, 4, 61
	if a < b {
		fmt.Println("b is greater")
	} else if b < c {
		fmt.Println("c is greater")
	} else {
		fmt.Println("a is grater than other")
	}
}

func useIfandShortHand() {

	if err := doSomething(); err != nil {
		fmt.Println("shortHand ")
	}
}
func doSomething() error {
	return nil
}
