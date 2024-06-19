package main

import "fmt"

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}

	// We can select other pivot as well
	pivot := arr[(len(arr)-1)/2]
	var less, greater []int
	pivots := []int{pivot}

	for i := range arr {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			greater = append(greater, arr[i])
		} else if i != (len(arr)-1)/2 {
			//This block takes care of repeating values
			pivots = append(pivots, arr[i])
		}
	}
	return append(append(qSort(less), pivots...), qSort(greater)...)
}

func main() {
	a := []int{6, 6, 1, 2, 5, 4, 3, -2, 0, 6}
	fmt.Println("before qSort:", a)
	fmt.Println("after qSort:", qSort(a))
}
