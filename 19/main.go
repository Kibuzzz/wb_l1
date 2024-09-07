package main

import (
	"fmt"
	"slices"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseString(s string) string {
	// Преобразуем строку в срез рун для корректной обработки Unicode символов.
	runes := []rune(s)
	// Используем стандартную библиотеку slices для переворота среза рун.
	slices.Reverse(runes)
	reversedString := string(runes)
	return reversedString
}

func main() {
	fmt.Println(reverseString("12345678910"))
	fmt.Println(reverseString("абвгдежзи"))
}
