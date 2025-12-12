// Go поддерживает
// <a href="https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D0%BA%D1%83%D1%80%D1%81%D0%B8%D1%8F"><em>рекурсивные функции</em></a>.
// Вот классический пример.

package main

import "fmt"

// Эта функция `fact` вызывает сама себя до тех пор,
// пока не достигнет базового случая `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Анонимные функции тоже могут быть рекурсивными, но для
	// этого нужно явно объявить переменную через `var` для
	// хранения функции до её определения.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Поскольку `fib` была объявлена ранее в `main`,
		// Go знает, какую функцию вызывать через `fib`.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
