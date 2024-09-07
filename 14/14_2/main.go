package main

import (
	"fmt"
)

func detectType(a interface{}) string {
	switch a.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan"
	default:
		return "неизвестно"
	}
}

func main() {
	var i int
	var s string
	var b bool
	var c chan int
	fmt.Println(detectType(i))
	fmt.Println(detectType(s))
	fmt.Println(detectType(b))
	fmt.Println(detectType(c))
}
