package main

import (
	"fmt"
	"strings"
)

/* var justString string
 * func someFunc() {
 *  v := createHugeString(1 << 10)
 *  justString = v[:100]
 * }
 *
 * func main() {
 *  someFunc()
 * }
 *
 * v isn't deleted after someFunc()
 * also v[:100] may lead to unwanted behavior
 * with unicode
 */

func createHugeString(size int) string {
	// Use string builder to create big unicode string
	var sb strings.Builder
	for i := 0; i < size; i++ {
		sb.WriteString("好")
	}
	return sb.String()
}

var justString string

func someFunc() {
	var v []rune
	// Slice from []rune will work as expected with unicode
	v = []rune(createHugeString(1 << 10))
	// Deep copy allows v to be released
	justString := strings.Clone(string(v[:100]))
	fmt.Println(justString)
}

func main() {
	someFunc()
}
