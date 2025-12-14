// Распространённая задача в программах — получить количество
// секунд, миллисекунд или наносекунд с момента
// [эпохи Unix](https://en.wikipedia.org/wiki/Unix_time).
// Вот как это сделать в Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Используй `time.Now` с `Unix`, `UnixMilli` или `UnixNano`,
	// чтобы получить прошедшее время с эпохи Unix в секундах,
	// миллисекундах или наносекундах соответственно.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// Можно также преобразовать целые секунды или наносекунды
	// с эпохи в соответствующее значение `time`.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
