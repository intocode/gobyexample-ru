// Стандартная библиотека Go поставляется с отличной
// поддержкой HTTP-клиентов и серверов в пакете `net/http`.
// В этом примере мы используем его для выполнения
// простых HTTP-запросов.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Выполняем HTTP GET запрос к серверу. `http.Get` —
	// удобное сокращение для создания объекта `http.Client`
	// и вызова его метода `Get`; он использует объект
	// `http.DefaultClient` с полезными настройками
	// по умолчанию.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Выводим статус HTTP-ответа.
	fmt.Println("Response status:", resp.Status)

	// Выводим первые 5 строк тела ответа.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
