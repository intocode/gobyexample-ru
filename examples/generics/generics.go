// Начиная с версии 1.18, в Go добавлена поддержка
// _дженериков_, также известных как _параметры типов_.

package main

import "fmt"

// В качестве примера обобщённой функции `SlicesIndex` принимает
// слайс любого `comparable` типа и элемент этого типа,
// возвращая индекс первого вхождения v в s, или -1, если
// элемент отсутствует. Ограничение `comparable` означает,
// что мы можем сравнивать значения этого типа операторами
// `==` и `!=`. Подробное объяснение этой сигнатуры типа
// см. в [этой статье](https://go.dev/blog/deconstructing-type-parameters).
// Обрати внимание, что эта функция существует в стандартной
// библиотеке как [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// В качестве примера обобщённого типа `List` — это
// односвязный список со значениями любого типа.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Мы можем определять методы для обобщённых типов так же,
// как и для обычных, но нужно сохранять параметры типов.
// Тип — это `List[T]`, а не `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements возвращает все элементы List в виде слайса.
// В следующем примере мы увидим более идиоматичный способ
// итерации по всем элементам пользовательских типов.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// При вызове обобщённых функций часто можно положиться
	// на _вывод типов_. Обрати внимание, что нам не нужно
	// указывать типы для `S` и `E` при вызове `SlicesIndex` —
	// компилятор выводит их автоматически.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ...хотя мы могли бы указать их явно.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
