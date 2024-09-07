package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseString(s string) string {
	// Преобразуем строку в срез рун для корректной обработки Unicode символов.
	runes := []rune(s)
	// Используем стандартную библиотеку slices для переворота среза рун.
	slices.Reverse(runes)
	reversedString := string(runes)
	return reversedString
}

func reverseWords(s string) string {
	// Разделяю строку на слова с помощью стандартной библиотеки
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = reverseString(word)
	}
	reversedWords := strings.Join(words, " ")
	return reversedWords
}

func main() {
	fmt.Println(reverseWords("snow dog sun — sun dog snow")) // wons god nus — nus god wons
}
