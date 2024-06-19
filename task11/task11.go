package main

import "fmt"

func intersection(x []int, y []int) []int {
	// map is used to count how many times number is present in x and y
	m := make(map[int]int)
	var res []int
	for _, v := range x {
		m[v] += 1
	}
	for _, v := range y {
		m[v] += 1
	}
	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}
	return res
}
func main() {
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(intersection(a1, a2))
}
