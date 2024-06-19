package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	// Case insensitive, so A=a, B=b, etc.
	lows := strings.ToLower(s)
	// Use map with count to check uniqueness
	countMap := map[rune]int{}
	for _, v := range lows {
		countMap[v]++
	}
	for _, v := range countMap {
		if v > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("abcd:", checkUnique("abcd"))
	fmt.Println("abCdefAaf:", checkUnique("abCdefAaf"))
	fmt.Println("aabcd:", checkUnique("aabcd"))
}
