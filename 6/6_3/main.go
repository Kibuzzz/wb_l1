package main

import (
	"fmt"
	"sync"
	"time"
)

// Способ 3: горутина останавливается, когда приходит сигнал из канала
func main() {

	cancelCh := make(chan struct{})

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-cancelCh:
				fmt.Println("закрытие горутины: пришел сигнал из канала")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("полезная работа")
			}
		}
	}()
	time.Sleep(time.Second * 5)
	cancelCh <- struct{}{}
	wg.Wait() // ждем завершения горутины
	fmt.Println("завершение программы")
}
