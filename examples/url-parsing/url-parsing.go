// URL предоставляют [унифицированный способ адресации ресурсов](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Вот как парсить URL в Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Распарсим этот пример URL, который включает схему,
	// данные аутентификации, хост, порт, путь, параметры
	// запроса и фрагмент.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Парсим URL и проверяем на отсутствие ошибок.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Доступ к схеме прост.
	fmt.Println(u.Scheme)

	// `User` содержит всю информацию для аутентификации;
	// вызови `Username` и `Password` для получения
	// отдельных значений.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` содержит и имя хоста, и порт (если он есть).
	// Используй `SplitHostPort` для их извлечения.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Здесь мы извлекаем `path` и фрагмент после `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Чтобы получить параметры запроса в строковом формате
	// `k=v`, используй `RawQuery`. Можно также распарсить
	// параметры запроса в карту. Карты параметров — это
	// соответствия строк срезам строк, поэтому используй
	// индекс `[0]`, если нужно только первое значение.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
