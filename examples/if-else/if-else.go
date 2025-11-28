// В Go ветвление с помощью `if` и `else` достаточно простое.

package main

import "fmt"

func main() {

	// Вот простой пример.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// У `if` может не быть ветки `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// В условиях часто используются логические операторы вроде `&&` и `||`.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Перед условием в `if` можно писать выражения; любые
	// переменные, объявленные в нём, будут доступны в текущем `if`
	// и всех последующих ветках (то есть в связанных `else`).
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Обрати внимание: в Go вокруг условия не нужны скобки,
// но фигурные скобки обязательны.
