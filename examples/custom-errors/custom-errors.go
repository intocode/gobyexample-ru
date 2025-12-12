// Можно определять пользовательские типы ошибок,
// реализовав на них метод `Error()`. Вот вариант
// примера выше, который использует пользовательский тип
// для явного представления ошибки аргумента.

package main

import (
	"errors"
	"fmt"
)

// Пользовательский тип ошибки обычно имеет суффикс "Error".
type argError struct {
	arg     int
	message string
}

// Добавление этого метода `Error` делает `argError`
// реализацией интерфейса `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Возвращаем нашу пользовательскую ошибку.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.As` — это более продвинутая версия `errors.Is`.
	// Она проверяет, соответствует ли данная ошибка (или любая
	// ошибка в её цепочке) определённому типу ошибки, и преобразует
	// её в значение этого типа, возвращая `true`. Если совпадения
	// нет, возвращается `false`.
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
