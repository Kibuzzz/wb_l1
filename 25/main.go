package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func mySleep(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	<-ctx.Done()
	cancel()
}

func main() {
	t := time.Now()
	mySleep(time.Second * 2)
	fmt.Printf("Времени прошло: %s\n", time.Since(t))
}
