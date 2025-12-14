// Пакет `strings` из стандартной библиотеки предоставляет
// множество полезных функций для работы со строками. Вот
// несколько примеров, чтобы дать представление о пакете.

package main

import (
	"fmt"
	s "strings"
)

// Создаём псевдоним для `fmt.Println` с коротким именем,
// так как будем часто его использовать.
var p = fmt.Println

func main() {

	// Вот примеры функций, доступных в пакете `strings`.
	// Поскольку это функции пакета, а не методы самого
	// строкового объекта, нужно передавать строку первым
	// аргументом функции. Больше функций можно найти в
	// документации пакета [`strings`](https://pkg.go.dev/strings).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
