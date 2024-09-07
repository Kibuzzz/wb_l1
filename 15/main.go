package main

import (
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

var justString string

// Проблема в том, что justString будет ссылаться лишь часть большой строки, но в памяти будет
// хранится вся большая строка, из за этого большое количество памяти будет занято неиспользуемыми данными

// Чтобы этого избежать можно создать новую строку, а не ссылку на оригинальную.
// Так большая строка будет собрана сборщиком мусора и не будет занимать место.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
	_ = justString
}
