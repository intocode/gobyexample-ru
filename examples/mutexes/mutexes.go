// В предыдущем примере мы рассмотрели управление простым
// состоянием счётчика с помощью [атомарных операций](atomic-counters).
// Для более сложного состояния можно использовать [_мьютекс_](https://en.wikipedia.org/wiki/Mutual_exclusion),
// чтобы безопасно обращаться к данным из нескольких горутин.

package main

import (
	"fmt"
	"sync"
)

// Container содержит словарь счётчиков; поскольку мы хотим
// обновлять его конкурентно из нескольких горутин, добавляем
// `Mutex` для синхронизации доступа.
// Обрати внимание, что мьютексы нельзя копировать, поэтому
// если эта структура передаётся куда-либо, это нужно делать
// по указателю.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Блокируем мьютекс перед доступом к `counters`; разблокируем
	// его в конце функции с помощью оператора [defer](defer).
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Обрати внимание, что нулевое значение мьютекса готово к
		// использованию, поэтому инициализация здесь не требуется.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Эта функция увеличивает именованный счётчик
	// в цикле.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Запускаем несколько горутин конкурентно; обрати
	// внимание, что все они обращаются к одному `Container`,
	// а две из них — к одному и тому же счётчику.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Ждём завершения горутин
	wg.Wait()
	fmt.Println(c.counters)
}
