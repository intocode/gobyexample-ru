// [Переменные окружения](https://en.wikipedia.org/wiki/Environment_variable) —
// универсальный механизм для [передачи конфигурации
// Unix-программам](https://www.12factor.net/config).
// Рассмотрим, как устанавливать, получать и выводить
// переменные окружения.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Чтобы установить пару ключ/значение, используй
	// `os.Setenv`. Чтобы получить значение по ключу,
	// используй `os.Getenv`. Вернётся пустая строка,
	// если ключа нет в окружении.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Используй `os.Environ` для получения списка всех
	// пар ключ/значение в окружении. Возвращается срез
	// строк вида `KEY=value`. Можно использовать
	// `strings.SplitN` для получения ключа и значения.
	// Здесь мы выводим все ключи.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
