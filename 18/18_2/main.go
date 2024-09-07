package main

import (
	"fmt"
	"sync"
)

type counter struct {
	n  int64
	mu *sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func increment(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c.increment()
	}
}

func main() {
	c := &counter{n: 0, mu: &sync.Mutex{}}

	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment(c, wg)
	}
	wg.Wait()
	fmt.Println(c.n)
}
