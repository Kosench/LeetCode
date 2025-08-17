package __Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int) // хранит последний индекс символа
	left := 0
	maxLen := 0

	for right, ch := range s {
		// если символ уже встречался и его индекс внутри текущего окна
		if prevIndex, ok := charIndex[ch]; ok && prevIndex >= left {
			left = prevIndex + 1
		}
		charIndex[ch] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
