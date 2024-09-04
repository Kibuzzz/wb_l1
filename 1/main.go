package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) GetName() string {
	return fmt.Sprintf("Привет, мое имя - %s", h.Name)
}

func (h *Human) Greet() string {
	return "Привет"
}

func (h *Human) UniqueMethod() string {
	return "Уникальный метод Human"
}

type Action struct {
	Name string
	Human
}

func (a *Action) GetName() string {
	return fmt.Sprintf("Название действия - %s", a.Name)
}

func (a *Action) Greet(name string) string {
	return fmt.Sprintf("Action: поприветствовать %s", name)
}

func main() {
	a := Action{Name: "Ходить"}
	fmt.Printf("%v\n", a) // Cтруктура Human с дефолтными значениями
	h := Human{Name: "Oleg", Age: 32}
	a.Human = h

	// Обращение к полям родительской и встроенной структуры
	fmt.Println(a.Name)       // Ходить
	fmt.Println(a.Human.Name) // Oleg

	// Если у родительской и встроенной структуры одинаковые методы,
	// тогда, чтобы обратиться к методу структуры, нужно явно указать,
	// что вызывается метод у встроенной структуры
	fmt.Println(a.GetName())       // Название действия - Ходить
	fmt.Println(a.Human.GetName()) // Привет, мое имя - Oleg

	// Если имя поля/метода встроенной структуры уникальное,
	// то можно обратиться к нему напрямую
	fmt.Println(a.Age)            // 32
	fmt.Println(a.UniqueMethod()) // Уникальный метод Human

	// Можно переопределить метод встроенной структуры
	fmt.Println(h.Greet())        // Привет
	fmt.Println(a.Human.Greet())  // Привет
	fmt.Println(a.Greet("Игнат")) // Action: поприветствовать Игнат
}
