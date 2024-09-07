package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	n int64
}

func increment(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&c.n, 1)
	}
}

func main() {
	c := &counter{}

	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment(c, wg)
	}
	wg.Wait()
	fmt.Println(c.n)
}
