package main

import (
	"fmt"
	"sync"
	"time"
)

func workerTwo(input chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if _, ok := <-input; ok {
			time.Sleep(time.Second * 1)
			fmt.Println("полезная работа горутины 2")
		} else {
			fmt.Println("закрытие горутины: канал закрыт")
			return
		}
	}
}

func workerOne(input chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range input {
		time.Sleep(time.Second * 1)
		fmt.Println("полезная работа горутины 1")
	}
	fmt.Println("закрытие горутины 1: канал закрыт")
}

// Способ 1: горутина заканчивается, когда закрывается канал, из которого она читает
func main() {
	input := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go workerOne(input, wg)
	// go workerTwo(input, wg)

	for i := 0; i < 5; i++ { // исправленный цикл
		input <- i
	}
	close(input)
	fmt.Println("канал закрыт")
	wg.Wait() // ждем завершения горутины
}
