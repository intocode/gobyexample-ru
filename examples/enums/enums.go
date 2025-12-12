// _Перечисляемые типы_ (enum) — это частный случай
// [типов-сумм](https://ru.wikipedia.org/wiki/%D0%A2%D0%B8%D0%BF-%D1%81%D1%83%D0%BC%D0%BC%D0%B0).
// Enum — это тип с фиксированным набором возможных
// значений, каждое из которых имеет своё имя. В Go нет
// enum как отдельной языковой конструкции, но их легко
// реализовать с помощью существующих идиом языка.

package main

import "fmt"

// Наш enum-тип `ServerState` имеет базовый тип `int`.
type ServerState int

// Возможные значения для `ServerState` определены как
// константы. Специальное ключевое слово [iota](https://go.dev/ref/spec#Iota)
// автоматически генерирует последовательные значения
// констант; в данном случае 0, 1, 2 и так далее.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Реализация интерфейса [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// позволяет выводить значения `ServerState` на печать
// или преобразовывать их в строки.
//
// Это может быть громоздко при большом количестве значений.
// В таких случаях можно использовать инструмент
// [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// совместно с `go:generate` для автоматизации процесса.
// См. [эту статью](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// для подробного объяснения.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Если у нас есть значение типа `int`, мы не можем передать
	// его в `transition` — компилятор сообщит о несоответствии типов.
	// Это обеспечивает некоторую степень типобезопасности enum
	// на этапе компиляции.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition эмулирует переход состояния сервера;
// принимает текущее состояние и возвращает новое.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Предположим, здесь мы проверяем некоторые
		// условия для определения следующего состояния...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
