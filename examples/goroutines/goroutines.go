// _Goroutine_ — это легковесный поток выполнения.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Допустим, у нас есть вызов функции `f(s)`. Вот как
	// мы бы вызвали её обычным способом, выполняя
	// синхронно.
	f("direct")

	// Чтобы вызвать эту функцию в goroutine, используй
	// `go f(s)`. Эта новая goroutine будет выполняться
	// конкурентно с вызывающей.
	go f("goroutine")

	// Также можно запустить goroutine для вызова
	// анонимной функции.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Оба наших вызова функций теперь выполняются
	// асинхронно в отдельных goroutine. Подождём их
	// завершения (для более надёжного подхода
	// используй [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
