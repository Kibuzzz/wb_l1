package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := []string{}

	// Будет хранить в себе уже встречавшиеся слова
	dictionary := make(map[string]bool)
	for _, word := range words {
		exists := dictionary[word]
		// Если слово уже было записано в словарь, знать сейчас дубликат, его добавлять не нужно
		if !exists {
			set = append(set, word)
			dictionary[word] = true
		}
	}
	fmt.Println(set)
}
