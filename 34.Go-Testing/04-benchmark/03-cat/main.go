package main

import (
	"fmt"
	"strings"
	"test-branch/mystr"
)

const s = "Dear Arjun,As we're approaching the final stages of the recruiting process, we need some information from you to work on the offer release process.Request you to fill up this form at the earliest. Below is the form link."

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))

}
