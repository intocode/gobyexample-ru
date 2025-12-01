// Go поддерживает [_анонимные функции_](https://en.wikipedia.org/wiki/Anonymous_function),
// которые могут формировать <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>замыкания</em></a>.
// Анонимные функции полезны, когда вы хотите объявить
// строчную функцию, не давая ей определённого имени.

package main

import "fmt"

// Эта функция `intSeq` возвращает другую функцию, которую
// мы объявили анонимно в теле `intSeq`. Возвращённая функция
// _замыкает_ переменную `i` для формирования
// замыкания.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Мы вызываем `intSeq`, сохраняя результат (функцию)
	// в переменную `nextInt`. Значение текущей функции перехватывает
	// значение внутренней `i`, которое будет обновлено
	// каждый раз при вызове `nextInt`.
	nextInt := intSeq()

	// Смотрим на эффект замыкания, вызывая несколько раз `nextInt`.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Для подтверждения что состояние уникально текущей функции,
	// создаём и тестируем новую функцию.
	newInts := intSeq()
	fmt.Println(newInts())
}
