package main

import "fmt"

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			// Try skipping one character from either left or right
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

// Helper function to check if substring is palindrome
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "absa"
	fmt.Println(validPalindrome(s))
}
