package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func calculateSumSquares(nums []int) {
	var sum int64
	wg := &sync.WaitGroup{}
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			square := num * num
			atomic.AddInt64(&sum, int64(square))
		}(num)
	}
	wg.Wait()
	fmt.Println(sum)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	calculateSumSquares(nums)
}
