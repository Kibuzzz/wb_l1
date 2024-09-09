package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20

func sum(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func diff(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	// Инициализация больших чисел a и b
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("20971520", 10)
	b.SetString("10485760", 10)

	// Выполнение арифметических операций
	sum := sum(a, b)
	diff := diff(a, b)
	product := mul(a, b)
	div := div(a, b)

	// Вывод результатов
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("Сложение (a + b) = %v\n", sum)
	fmt.Printf("Вычитание (a - b) = %v\n", diff)
	fmt.Printf("Умножение (a * b) = %v\n", product)
	fmt.Printf("Деление (a / b) = %v\n", div) //
}
