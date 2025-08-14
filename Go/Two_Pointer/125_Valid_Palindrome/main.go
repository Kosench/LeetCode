package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}
