package main

import "fmt"

// Удалить i-ый элемент из слайса.

func deleteI(nums []int, i int) []int {
	if i < 0 || i > len(nums) {
		return nums
	}
	nums = append(nums[:i], nums[i+1:]...)
	return nums
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(deleteI(nums, 5))
}
