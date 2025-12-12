// Мы можем использовать каналы для синхронизации
// выполнения между goroutine. Вот пример использования
// блокирующего получения для ожидания завершения
// goroutine. При ожидании завершения нескольких
// goroutine лучше использовать [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// Это функция, которую мы запустим в goroutine. Канал
// `done` будет использоваться для уведомления другой
// goroutine о завершении работы этой функции.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Отправляем значение, чтобы уведомить о завершении.
	done <- true
}

func main() {

	// Запускаем worker goroutine, передавая ей канал
	// для уведомления.
	done := make(chan bool, 1)
	go worker(done)

	// Блокируемся, пока не получим уведомление от
	// worker через канал.
	<-done
}
