package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором

type Point struct {
	x float64
	y float64
}

func New(x, y float64) Point {
	return Point{x: x, y: y}
}

func distance(p1, p2 Point) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	dist := math.Sqrt(x*x + y*y)
	return dist
}

func main() {
	a := New(1, 1)
	b := New(4, 4)
	fmt.Println(distance(a, b))
}
