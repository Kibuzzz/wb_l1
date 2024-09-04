package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Способ 2: горутина заканчивается, когда завершается контекст
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("закрытие горутины: контекс завершен")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("полезная работа")
			}
		}
	}()
	wg.Wait() // ждем завершения горутины
	fmt.Println("завершение программы")
}
