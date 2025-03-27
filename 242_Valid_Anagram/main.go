package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runes := []rune(s)
	runes2 := []rune(t)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(runes) == string(runes2)
}

func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := [26]int{}

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}

func main() {
	str := "anagram"
	str_2 := "nagaram"

	fmt.Println(isAnagram(str, str_2))
}
