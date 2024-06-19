package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Struct constructor
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func findDistance(point1, point2 *Point) float64 {
	return math.Sqrt(math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2))
}

func main() {
	a := NewPoint(1, 1)
	b := NewPoint(0, 0)
	fmt.Println("Distance:", findDistance(a, b))
}
