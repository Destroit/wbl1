package main

import "fmt"

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32)
	for _, v := range temperature {
		group := int(v/10) * 10
		m[group] = append(m[group], v)
	}
	fmt.Println(m)
}
