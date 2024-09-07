package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func checkDuplicates(s string) bool {
	// Чтобы функция была регистронезависимой, приводим все буквы к нижнему регистру
	s = strings.ToLower(s)
	dict := make(map[rune]struct{})
	for _, rune := range s {
		_, duplicate := dict[rune]
		if duplicate {
			return false
		}
		dict[rune] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(checkDuplicates("abcd"))
	fmt.Println(checkDuplicates("abCdefAaf"))
	fmt.Println(checkDuplicates("aabcd"))
}
