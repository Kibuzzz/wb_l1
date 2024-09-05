package main

import "fmt"

func main() {
	nums1 := []int{2, 4, 6, 8, 10}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	intersectionMap := make(map[int]bool)
	intersectons := []int{}
	// Заполняем мапу, где дальше будет проверяться повторение
	for _, num := range nums1 {
		intersectionMap[num] = true
	}
	for _, num := range nums2 {
		// проверяем лежит ли такое число уже в мапе
		_, ok := intersectionMap[num]
		// если нашли повторение, добавляем в слайс пересечения
		if ok {
			intersectons = append(intersectons, num)
		}
	}
	fmt.Println(intersectons)
}
