package main

import "fmt"

func main() {
	x := 10
	y := 5
	fmt.Printf("x: %d, y: %d\n", x, y)
	// Гибридная форма x & y
	x = x ^ y
	// Вычитаем из "гибридной формы" y, остается только х
	y = x ^ y
	// Теперь y = x, вычитаем из "гибридной формы" информацию об х
	// остается y
	x = x ^ y
	fmt.Printf("x: %d, y: %d\n", x, y)
}
