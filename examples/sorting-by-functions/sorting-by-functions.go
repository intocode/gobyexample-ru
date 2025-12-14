// Иногда нужно отсортировать коллекцию не в естественном
// порядке. Например, мы хотим отсортировать строки по длине,
// а не по алфавиту. Вот пример пользовательской сортировки
// в Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Реализуем функцию сравнения для длин строк.
	// Для этого удобно использовать `cmp.Compare`.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Теперь можно вызвать `slices.SortFunc` с этой
	// пользовательской функцией сравнения, чтобы
	// отсортировать `fruits` по длине названий.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// Тот же подход можно использовать для сортировки
	// среза значений, которые не являются встроенными типами.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// Сортируем `people` по возрасту с помощью `slices.SortFunc`.
	//
	// Примечание: если структура `Person` большая,
	// возможно, стоит использовать срез `*Person`
	// и соответствующим образом изменить функцию сортировки.
	// Если сомневаетесь — [бенчмарк](testing-and-benchmarking)!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
