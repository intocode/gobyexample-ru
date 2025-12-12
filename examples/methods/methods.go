// Go поддерживает _методы_, определённые для типов структур.

package main

import "fmt"

type rect struct {
	width, height int
}

// Этот метод `area` имеет _тип получателя_ `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Методы могут быть определены как для указателей, так и для
// значений. Вот пример получателя по значению.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Здесь мы вызываем два метода, определённых для структуры.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go автоматически преобразует значения и указатели
	// при вызове методов. Получатель-указатель позволяет
	// избежать копирования при вызове метода или даёт
	// возможность изменять получающую структуру.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
