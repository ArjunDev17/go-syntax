package main

import (
	"fmt"
	"strings"
)

func countVowel(st string) int {
	vowel := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}
	var count int
	st = strings.ToLower(st)
	for _, ch := range st {

		if vowel[ch] {

			count++
		}
		fmt.Printf("%c", ch)
	}
	return count
}
func sunStr(s string, k int) int {
	maxCount := 0
	for i := 0; i < len(s); i += k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		fmt.Println(s[i:end], "\n")
		count := countVowel(s[i:end])
		if maxCount < count {
			maxCount = count
		}

	}
	return maxCount

}
func maxVowels(s string, k int) int {
	// Define vowels
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	// Count vowels in the first window of size k
	count := 0
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}

	// Track the maximum number of vowels found
	maxCount := count

	// Slide the window from position 1 to len(s)-k
	for i := k; i < len(s); i++ {
		// Remove the effect of the character that is sliding out of the window
		if vowels[s[i-k]] {
			count--
		}
		// Add the effect of the new character that is sliding into the window
		if vowels[s[i]] {
			count++
		}
		// Update maxCount if the new count is higher
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func main() {
	s := "weallloveyou"
	k := 7

	res := maxVowels(s, k)
	fmt.Println("\nres :", res)
}
