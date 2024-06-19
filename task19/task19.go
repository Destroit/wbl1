package main

import (
	"fmt"
)

func reverse(s string) string {
	// Rune for unicode
	r := []rune(s)
	// Swapping from edges to center
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main() {
	a := "你好"
	fmt.Println(a)
	fmt.Println("Reverse: ", reverse(a))
	b := "абырвалг"
	fmt.Println(b)
	fmt.Println("Reverse: ", reverse(b))
	c := "😀😐💯"
	fmt.Println(c)
	fmt.Println("Reverse: ", reverse(c))
}
