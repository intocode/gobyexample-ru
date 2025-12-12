// _Интерфейсы_ — это именованные коллекции сигнатур
// методов.

package main

import (
	"fmt"
	"math"
)

// Вот базовый интерфейс для геометрических фигур.
type geometry interface {
	area() float64
	perim() float64
}

// Для примера мы реализуем этот интерфейс для
// типов `rect` и `circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Чтобы реализовать интерфейс в Go, нужно просто
// реализовать все методы этого интерфейса. Здесь мы
// реализуем `geometry` для `rect`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Реализация для `circle`.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Если переменная имеет тип интерфейса, мы можем вызывать
// методы, входящие в этот интерфейс. Вот обобщённая
// функция `measure`, которая использует это для работы
// с любой `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Иногда полезно узнать тип значения интерфейса во время
// выполнения. Один из способов — использовать *утверждение
// типа*, как показано здесь; другой — [type `switch`](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Типы структур `circle` и `rect` оба реализуют
	// интерфейс `geometry`, поэтому мы можем использовать
	// экземпляры этих структур в качестве аргументов
	// для `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
