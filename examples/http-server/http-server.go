// Написать базовый HTTP-сервер легко с использованием
// пакета `net/http`.
package main

import (
	"fmt"
	"net/http"
)

// Фундаментальная концепция серверов `net/http` —
// *обработчики*. Обработчик — это объект, реализующий
// интерфейс `http.Handler`. Распространённый способ
// написания обработчика — использование адаптера
// `http.HandlerFunc` на функциях с подходящей сигнатурой.
func hello(w http.ResponseWriter, req *http.Request) {

	// Функции, служащие обработчиками, принимают
	// `http.ResponseWriter` и `http.Request` как аргументы.
	// Response writer используется для заполнения
	// HTTP-ответа. Здесь наш простой ответ — просто
	// "привет\n".
	fmt.Fprintf(w, "привет\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Этот обработчик делает кое-что посложнее: читает
	// все HTTP-заголовки запроса и возвращает их
	// в теле ответа.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Регистрируем наши обработчики на маршрутах сервера
	// с помощью удобной функции `http.HandleFunc`. Она
	// настраивает *роутер по умолчанию* в пакете `net/http`
	// и принимает функцию как аргумент.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Наконец, вызываем `ListenAndServe` с портом и
	// обработчиком. `nil` указывает использовать роутер
	// по умолчанию, который мы только что настроили.
	http.ListenAndServe(":8090", nil)
}
