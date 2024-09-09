package main

import (
	"fmt"
)

func quicksort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	pivot := len(nums) - 1 // Индекс опорного элемента (последний элемент массива)
	partitionIndex := 0    // Индекс границы для разделения элементов, которые меньше опорного элемента

	for i := range nums {
		// Если текущий элемент меньше опорного, перемещаем его влево от границы
		if nums[i] < nums[pivot] {
			nums[i], nums[partitionIndex] = nums[partitionIndex], nums[i]
			partitionIndex++
		}
	}

	// Помещаем опорный элемент в его окончательную позицию
	nums[partitionIndex], nums[pivot] = nums[pivot], nums[partitionIndex]

	// Рекурсивно применяем быструю сортировку к левой и правой частям
	quicksort(nums[:partitionIndex])   // Сортируем элементы слева от опорного
	quicksort(nums[partitionIndex+1:]) // Сортируем элементы справа от опорного

	return nums
}

func main() {
	nums := []int{17, 142, 4, 200, 5, 6, 1, 1}
	fmt.Println(quicksort(nums))
}
