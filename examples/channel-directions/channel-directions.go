// При использовании каналов как параметров функции
// можно указать, предназначен ли канал только для
// отправки или только для получения значений. Эта
// специфичность повышает типобезопасность программы.

package main

import "fmt"

// Эта функция `ping` принимает канал только для отправки
// значений. Попытка получить из этого канала приведёт
// к ошибке компиляции.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Функция `pong` принимает один канал для получения
// (`pings`) и второй для отправки (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
