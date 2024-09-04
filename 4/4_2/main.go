package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func startWorkers(workerNum int, input chan any, wg *sync.WaitGroup) {
	for range workerNum {
		go worker(input, wg)
	}
}

func worker(input chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range input {
		fmt.Println(n)
	}
	fmt.Println("worker done")
}

type Test struct {
	Name string
}

func main() {
	var workerNum int

	fmt.Println("Введите количество воркеров:")
	fmt.Scanln(&workerNum)

	input := make(chan any)
	wg := &sync.WaitGroup{}
	wg.Add(workerNum)
	startWorkers(workerNum, input, wg)

	// Контекс с каналом Done, в который придет значение, когда будет вызван сигнал отмены
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Слайс с разными типами данных
	l := []any{"123", 5, Test{Name: "test struct"}}
	for _, a := range l {
		input <- a
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Got cancellation signal")
			close(input)
			wg.Wait()
			return
		default:
			time.Sleep(time.Second)
			n := rand.Intn(100)
			input <- n
		}
	}
}
