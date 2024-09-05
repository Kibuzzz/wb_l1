package main

import (
	"fmt"
	"log"
)

func main() {
	var num int = 42 // Пример числа
	var i int        // Бит, который нужно изменить
	fmt.Println("Введите бит который нужно изменить: ")
	_, err := fmt.Scanln(&i)
	if err != nil {
		log.Fatal("не получилось получить индекс бита")
	}
	if i < 0 {
		log.Fatal("число не может быть меньше 0")
	}

	fmt.Printf("Число до: %d, %b\n", num, num)
	num = num ^ (1 << i)
	fmt.Printf("Число после: %d, %b\n", num, num)
}
