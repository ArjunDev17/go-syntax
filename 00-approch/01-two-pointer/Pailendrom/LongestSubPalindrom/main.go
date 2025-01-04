package main

import "fmt"

func main() {
	str := "ababba"
	res := longestPailendrom(str)
	fmt.Println("Longest Palindrome:", res)
}

func longestPailendrom(str string) string {
	if len(str) == 0 {
		return ""
	}
	start, maxLength := 0, 1
	for i := 0; i < len(str); i++ {
		// Check for odd length palindrome (centered at str[i])
		left, right := i, i
		for left >= 0 && right < len(str) && str[left] == str[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}

		// Check for even length palindrome (centered between str[i] and str[i+1])
		left, right = i, i+1
		for left >= 0 && right < len(str) && str[left] == str[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}
	}
	return str[start : start+maxLength]
}
