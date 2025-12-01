// Go поддерживает
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>рекурсивные функции</em></a>.
// Вот классический пример.

package main

import "fmt"

// Эта функция `fact` вызывает саму себя, пока
// не дойдёт до конечного сценария `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Анонимные функции могут быть также рекурсивными, но они требуют
	// явного объявления переменной с `var` для хранения
	// функции до её объявления.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Учитывая, что `fib` была ранее объявлена в `main`, Go
		// знает к какой функции обратиться для вызова `fib` здесь.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
