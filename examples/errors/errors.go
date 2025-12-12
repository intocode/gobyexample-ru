// В Go идиоматично передавать ошибки через явное,
// отдельное возвращаемое значение. Это отличается от
// исключений в языках вроде Java, Python и Ruby,
// а также от перегруженного единственного значения
// результат/ошибка, которое иногда используется в C.
// Подход Go позволяет легко видеть, какие функции
// возвращают ошибки, и обрабатывать их с помощью тех же
// языковых конструкций, что и для других задач.
//
// Подробнее см. в документации [пакета errors](https://pkg.go.dev/errors)
// и в [этой статье в блоге](https://go.dev/blog/go1.13-errors).

package main

import (
	"errors"
	"fmt"
)

// По соглашению ошибки идут последним возвращаемым
// значением и имеют тип `error` — встроенный интерфейс.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` создаёт базовое значение `error`
		// с заданным сообщением об ошибке.
		return -1, errors.New("can't work with 42")
	}

	// Значение `nil` в позиции ошибки означает,
	// что ошибки не было.
	return arg + 3, nil
}

// Sentinel-ошибка — это заранее объявленная переменная,
// используемая для обозначения определённого состояния ошибки.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// Мы можем оборачивать ошибки в ошибки более
		// высокого уровня для добавления контекста.
		// Самый простой способ — использовать глагол
		// `%w` в `fmt.Errorf`. Обёрнутые ошибки образуют
		// логическую цепочку (A оборачивает B, которая
		// оборачивает C и т.д.), которую можно исследовать
		// с помощью функций вроде `errors.Is` и `errors.As`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// Идиоматично использовать встроенную проверку ошибки
		// в строке с `if`.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` проверяет, соответствует ли данная ошибка
			// (или любая ошибка в её цепочке) конкретному значению
			// ошибки. Это особенно полезно для обёрнутых или вложенных
			// ошибок, позволяя идентифицировать определённые типы
			// ошибок или sentinel-ошибки в цепочке ошибок.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
