package main

import "fmt"

func makeSet(s []string) []string {
	// map keys are unique values we're looking for
	m := make(map[string]int)
	var res []string
	for _, v := range s {
		m[v]++
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(makeSet(s))
}
