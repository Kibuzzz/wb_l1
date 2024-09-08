package main

import "fmt"

func binarySearch(target int, nums []int) bool {

	low := 0
	high := len(nums) - 1

	for low <= high {
		median := (low + high) / 2

		if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(nums) || nums[low] != target {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(9, items))  // true
	fmt.Println(binarySearch(11, items)) // false
}
