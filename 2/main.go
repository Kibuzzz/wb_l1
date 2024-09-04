package main

import (
	"fmt"
	"sync"
)

func printSquare(num int) {
	fmt.Printf("Квадрат числа %d - %d\n", num, num*num)
}

// Количество горутин = количество чисел
func calculateSquares(nums []int) {
	wg := &sync.WaitGroup{}
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			printSquare(num)
		}(num)
	}
	wg.Wait()
}

const workersNum = 2

func startWorkers(input <-chan int, wg *sync.WaitGroup) {
	for range workersNum {
		go worker(input, wg)
	}
}

func worker(input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Воркеры читают из канала и печатают квадрат числа, пока канал не закрыт
	for num := range input {
		printSquare(num)
	}
}

func calculateSquaresWorkers(nums []int) {
	// Канал из которого будут читать воркеры
	input := make(chan int, workersNum)

	// WaitGroup чтобы дождаться выполнения горутин
	wg := &sync.WaitGroup{}
	wg.Add(workersNum)

	startWorkers(input, wg)

	// Запись чисел в канал, из которого будут читать воркеры
	for _, num := range nums {
		input <- num
	}

	// Закрываем канал, чтобы воркеры остановились читать из канала и завершились
	close(input)
	wg.Wait()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	calculateSquares(nums)
	calculateSquaresWorkers(nums)
}
