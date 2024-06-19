package main

import (
	"fmt"
	"strings"
)

func reverseOrder(s string) string {
	splitStr := strings.Split(s, " ")

	var sb strings.Builder
	sb.WriteString(splitStr[len(splitStr)-1])

	for i := len(splitStr) - 2; i >= 0; i-- {
		sb.WriteString(" ")
		sb.WriteString(splitStr[i])
	}
	return sb.String()
}

func main() {
	fmt.Println(reverseOrder("snow dog sun"))
}
