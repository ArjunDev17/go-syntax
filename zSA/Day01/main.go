package main

/*
input:=[]int{2,4,5,6} target 6
output->[0,1]

*/
import "fmt"

func main() {
	fmt.Println("DSA 01")
	arr := []int{1, 2, 3, 4} //Target :-6
	Target := 7
	res := checkIndex(arr, Target)
	fmt.Println("res :", res)

}
func checkIndex(value []int, Target int) []int {
	res := []int{}

	for index := 0; index < len(value)-1; index++ {

		if (value[index] + value[index+1]) == Target {
			// fmt.Println("index data is :", value[index], value[index+1])
			// fmt.Println("Target = ->:", value[index]+value[index+1])

			res = append(res, index)
			res = append(res, index+1)
		}
	}
	return res
}
