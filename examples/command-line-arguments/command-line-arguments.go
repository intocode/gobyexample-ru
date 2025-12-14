// [_Аргументы командной строки_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments) —
// распространённый способ параметризации выполнения программ.
// Например, `go run hello.go` использует аргументы `run`
// и `hello.go` для программы `go`.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` предоставляет доступ к необработанным
	// аргументам командной строки. Обрати внимание, что
	// первое значение в этом срезе — путь к программе,
	// а `os.Args[1:]` содержит аргументы программы.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Можно получить отдельные аргументы обычной индексацией.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
