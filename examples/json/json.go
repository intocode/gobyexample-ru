// Go предоставляет встроенную поддержку кодирования и
// декодирования JSON, включая работу со встроенными
// и пользовательскими типами данных.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Мы будем использовать эти две структуры для демонстрации
// кодирования и декодирования пользовательских типов ниже.
type response1 struct {
	Page   int
	Fruits []string
}

// В JSON будут кодироваться/декодироваться только экспортируемые поля.
// Поля должны начинаться с заглавной буквы, чтобы быть экспортируемыми.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Сначала рассмотрим кодирование базовых типов данных
	// в строки JSON. Вот несколько примеров для атомарных
	// значений.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// А вот примеры для срезов и карт, которые кодируются
	// в JSON-массивы и объекты, как и следовало ожидать.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Пакет JSON может автоматически кодировать твои
	// пользовательские типы данных. Он включит в закодированный
	// вывод только экспортируемые поля и по умолчанию будет
	// использовать их имена в качестве ключей JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Можно использовать теги в объявлениях полей структуры
	// для настройки имён ключей в закодированном JSON.
	// Посмотри определение `response2` выше, чтобы увидеть
	// пример таких тегов.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Теперь рассмотрим декодирование JSON-данных в значения Go.
	// Вот пример для обобщённой структуры данных.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Нужно предоставить переменную, куда пакет JSON
	// сможет поместить декодированные данные. Этот
	// `map[string]interface{}` будет хранить карту строк
	// к произвольным типам данных.
	var dat map[string]interface{}

	// Вот само декодирование и проверка на связанные ошибки.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Чтобы использовать значения в декодированной карте,
	// нужно преобразовать их к соответствующему типу.
	// Например, здесь мы преобразуем значение в `num`
	// к ожидаемому типу `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Для доступа к вложенным данным требуется серия
	// преобразований.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Можно также декодировать JSON в пользовательские типы данных.
	// Это даёт преимущества дополнительной типобезопасности
	// в наших программах и устраняет необходимость утверждений типа
	// при доступе к декодированным данным.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// В примерах выше мы всегда использовали байты и строки
	// как промежуточное звено между данными и JSON-представлением
	// на стандартном выводе. Можно также потоково передавать
	// JSON-кодирование напрямую в `os.Writer`, например
	// `os.Stdout` или даже в тела HTTP-ответов.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// Потоковое чтение из `os.Reader`, например `os.Stdin`
	// или тел HTTP-запросов, выполняется с помощью `json.Decoder`.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
