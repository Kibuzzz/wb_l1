package main

import (
	"fmt"
	"sync"
)

func main() {
	first := make(chan int)
	second := make(chan int)

	// Чтобы дождаться завершения горутин
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// функция пишущая числа в первый канал
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			first <- i
		}
		// когда все числа записаны канал можно закрыть
		close(first)
	}()
	// читает из первого канала, умножает на 2 и передает во второй канал
	go func() {
		defer wg.Done()
		for num := range first {
			second <- num * 2
		}
		// когда все числа прочитаны(первый канал закрыт), второй канал закрывается
		// функция завершается
		close(second)
	}()
	for doubledNum := range second {
		fmt.Println(doubledNum)
	}
	// wg можно убрать, потому что чтение из второго канала прекратится, когда все горутины уже будут завершены
	wg.Wait()
}
