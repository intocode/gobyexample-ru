// Начиная с версии 1.23, в Go добавлена поддержка
// [итераторов](https://go.dev/blog/range-functions),
// что позволяет использовать range практически с чем угодно!

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Вернёмся к типу `List` из
// [предыдущего примера](generics). В том примере
// у нас был метод `AllElements`, который возвращал слайс
// всех элементов списка. С итераторами Go мы можем
// сделать это лучше — как показано ниже.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All возвращает _итератор_, который в Go является функцией
// с [особой сигнатурой](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Функция-итератор принимает другую функцию в качестве
		// параметра, по соглашению называемую `yield` (но
		// имя может быть произвольным). Она вызывает `yield` для
		// каждого элемента, по которому мы хотим итерироваться,
		// и проверяет возвращаемое значение `yield` для
		// возможного досрочного завершения.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Итерация не требует базовой структуры данных
// и даже не обязана быть конечной! Вот функция,
// возвращающая итератор по числам Фибоначчи: она
// продолжает работать, пока `yield` возвращает `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Поскольку `List.All` возвращает итератор, мы можем
	// использовать его в обычном цикле `range`.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// В пакетах вроде [slices](https://pkg.go.dev/slices)
	// есть много полезных функций для работы с итераторами.
	// Например, `Collect` принимает любой итератор и собирает
	// все его значения в слайс.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// Когда цикл достигает `break` или досрочного return,
		// функция `yield`, переданная итератору, возвращает `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
