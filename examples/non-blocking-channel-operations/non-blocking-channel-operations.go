// Обычные отправки и получения из каналов блокирующие.
// Однако мы можем использовать `select` с веткой `default`,
// чтобы реализовать _неблокирующие_ отправки, получения
// и даже неблокирующие многовариантные `select`.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Вот неблокирующее получение. Если значение доступно
	// в канале `messages`, то `select` выберет ветку
	// `<-messages` с этим значением. Если нет —
	// немедленно выполнится ветка `default`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Неблокирующая отправка работает аналогично. Здесь `msg`
	// не может быть отправлено в канал `messages`, потому что
	// канал не буферизован и нет получателя.
	// Поэтому выбирается ветка `default`.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Мы можем использовать несколько веток `case` перед
	// `default`, чтобы реализовать многовариантный
	// неблокирующий select. Здесь мы пытаемся неблокирующе
	// получить данные из обоих каналов `messages` и `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
