package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// отправляет числа в канал, пока контест не завершится
func writer(ctx context.Context, nums chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("writing")
			nums <- i
			i++
		}
	}
}

// читает из канала, пока контекс не завершен
func reader(ctx context.Context, nums <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			num := <-nums
			fmt.Printf("reading %d\n", num)
		}
	}
}

func main() {
	var timeout int
	fmt.Println("Введите количество секунд:")
	fmt.Scanln(&timeout)

	nums := make(chan int)
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	wg.Add(2)
	go writer(ctx, nums, wg)
	go reader(ctx, nums, wg)

	wg.Wait()
}
