package main

import "fmt"

// This doesn't cause side effects
func removeDeep(arr []int, index int) []int {
	res := make([]int, 0)
	res = append(res, arr[:index]...)
	return append(res, arr[index+1:]...)
}

// This causes side effects
func removeShallow(arr []int, index int) []int {
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("nums:", nums)
	shallowNums := removeShallow(nums, 4)
	fmt.Println("shallowNums:", shallowNums)

	shallowNums[0] = 500
	fmt.Println("changed shallowNums:", shallowNums)
	// Changing shallowNums changes nums as well. Also operation of removal messed with elements of num
	fmt.Println("nums after shallowNums change:", nums)

	deepNums := removeDeep(nums, 0)
	fmt.Println("deepNums:", deepNums)

	deepNums[0] = 333
	fmt.Println("changed deepNums:", deepNums)
	// deepNums and removeDeep won't change nums
	fmt.Println("nums after deepNums change:", nums)

}
