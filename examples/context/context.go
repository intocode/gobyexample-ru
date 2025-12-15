// В предыдущем примере мы рассмотрели настройку простого
// [HTTP-сервера](http-server). HTTP-серверы полезны для
// демонстрации использования `context.Context` для
// управления отменой. `Context` переносит дедлайны,
// сигналы отмены и другие значения области запроса
// через границы API и горутины.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `context.Context` создаётся для каждого запроса
	// механизмом `net/http` и доступен через метод
	// `Context()`.
	ctx := req.Context()
	fmt.Println("сервер: обработчик hello запущен")
	defer fmt.Println("сервер: обработчик hello завершён")

	// Ждём несколько секунд перед отправкой ответа клиенту.
	// Это может имитировать работу, выполняемую сервером.
	// Во время работы следим за каналом `Done()` контекста
	// на предмет сигнала о необходимости отменить работу
	// и вернуться как можно скорее.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "привет\n")
	case <-ctx.Done():
		// Метод `Err()` контекста возвращает ошибку,
		// объясняющую, почему канал `Done()` был закрыт.
		err := ctx.Err()
		fmt.Println("сервер:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Как и раньше, регистрируем наш обработчик на маршруте
	// "/hello" и начинаем обслуживание.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
