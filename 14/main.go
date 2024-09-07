package main

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func detectType(a interface{}) string {
	t := reflect.TypeOf(a)
	switch t.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	}
	return "Тип не поддерживается"
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
