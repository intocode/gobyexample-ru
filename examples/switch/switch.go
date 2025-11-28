// С помощью оператора `switch` можно описывать условные
// конструкции с несколькими ветками.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Вот простой пример `switch`.
	i := 2
	fmt.Print("Запишем ", i, " как ")
	switch i {
	case 1:
		fmt.Println("один")
	case 2:
		fmt.Println("два")
	case 3:
		fmt.Println("три")
	}

	// Можно перечислить несколько выражений в одном `case`,
	// разделив их запятыми. В этом примере мы также
	// используем необязательный вариант `default`.
	switch time.Now().Weekday() { // день недели
	case time.Saturday, time.Sunday:
		fmt.Println("Сейчас выходной")
	default:
		fmt.Println("Сейчас будний день")
	}

	// Конструкция `switch` без выражения — это
	// альтернативный способ записать логику if/else. Здесь
	// мы также показываем, что в `case` можно использовать
	// не только константы.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Еще нет двенадцати")
	default:
		fmt.Println("Сейчас после полудня")
	}

	// Конструкция `type switch` сравнивает типы вместо
	// значений. Её можно использовать, чтобы узнать
	// конкретный тип значения интерфейса. В этом примере
	// переменная `t` внутри ветки будет иметь тип,
	// соответствующий этой ветке.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Я bool")
		case int:
			fmt.Println("Я int")
		default:
			fmt.Printf("Неизвестный тип %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
