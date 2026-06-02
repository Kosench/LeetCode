package main

import (
	"fmt"
	"unicode"
)

func isAlnum(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b))
}

func toLower(b byte) byte {
	return byte(unicode.ToLower(rune(b)))
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if !isAlnum(s[left]) {
			left++
			continue
		}

		if !isAlnum(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlnum(s[left]) {
			left++
			continue
		}

		if !isAlnum(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(validPalindrome("race a car"))
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(validPalindrome("abca"))
}

func validPalindrome_2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return isPalindrome_2(left+1, right, s) || isPalindrome_2(left, right-1, s)
		}
	}
	return true
}
func isPalindrome_2(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
