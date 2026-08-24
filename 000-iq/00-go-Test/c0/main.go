package main

import (
	"fmt"
	"math"
	"time"
)

type Shaper interface {
	Area() float64
}
type Rectangle struct {
	width, hight float64
}
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {

	return math.Pi * c.radius * c.radius
}
func (r Rectangle) Area() float64 {
	hight := r.hight
	width := r.width
	return hight * width
}

func main() {
	fmt.Println("Interface demo")
	// rectangl := Rectangle{width: 5, hight: 4}
	// circle := Circle{radius: 4}
	// var s Shaper
	// // var s Shaper
	// s = circle
	// fmt.Printf("area of circle is %f \n:", s.Area())

	// s = rectangl
	// fmt.Printf("\narea of rectangle is %f \n", s.Area())
	// PatterPrint()
	go IncrementValue()
	time.Sleep(1 * time.Second)

}

func IncrementValue() {
	// twice := 2
	n := 1
	for i := 0; i < 10; i++ {
		// res := math.Pow(float64(twice))
		// fmt.Println(" twice of number :", res)
		n++
	}
}

// ******* 7
//  ***** 5
//   ***	3
//   *	1

func PatterPrint() {
	for i := 4; i > 0; i-- {
		co := 1
		for j := 7; j > 0; j-- {
			if j%co == 0 {
				fmt.Print("*")
			}

		}
		co++
		fmt.Print("\n")
	}

}

// i-->7

// i := 2--->4-
// 1 second
