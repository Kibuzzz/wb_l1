package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// функция раз в 100 миллисекунд будет увеличивать случайное число в мапе на единицу
// завершается, когда кончается контекс
func addNumber(ctx context.Context, wg *sync.WaitGroup, m map[int]int, mu *sync.Mutex) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			time.Sleep(time.Millisecond * 100)
			n := rand.Intn(100)
			// lock чтобы избешать конкуретной записи в map (без лока будет ошибка fatal error: concurrent map writes)
			mu.Lock()
			m[n]++
			mu.Unlock()
		}
	}
}

func main() {
	m := make(map[int]int)
	mu := &sync.Mutex{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go addNumber(ctx, wg, m, mu)
	go addNumber(ctx, wg, m, mu)
	go addNumber(ctx, wg, m, mu)
	wg.Wait()
	fmt.Println(m, len(m))
}
