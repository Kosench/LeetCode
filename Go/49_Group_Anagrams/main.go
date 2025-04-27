package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
)

func signature(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func groupAnagrams_Map(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	groupAnagramsMap := make(map[string][]string, len(strs))
	for _, val := range strs {
		key := signature(strings.ToLower(val))
		groupAnagramsMap[key] = append(groupAnagramsMap[key], val)
	}

	result := [][]string{}
	for _, group := range groupAnagramsMap {
		result = append(result, group)
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		sortedWord := []rune(word)
		slices.Sort(sortedWord)
		groups[string(sortedWord)] = append(groups[string(sortedWord)], word)
	}
	return slices.Collect(maps.Values(groups))
}

func main() {
	strSl := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strSl))
}
