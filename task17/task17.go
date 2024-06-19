package main

import "fmt"

func binarySearch(arr []int, x int) int {
	var low, high, mid int
	low, high = 0, len(arr)-1
	for low <= high {
		mid = (high + low) / 2
		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			// mid is an index of x
			return mid
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	s := 5
	fmt.Printf("searching: %v index: %v\n", s, binarySearch(a, s))
}
