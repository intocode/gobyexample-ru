// Go поддерживает _встраивание_ структур и интерфейсов
// для более удобной _композиции_ типов.
// Не путай это с [`//go:embed`](embed-directive) — директивой Go,
// появившейся в версии 1.16+ для встраивания файлов и папок
// в бинарный файл приложения.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` _встраивает_ `base`. Встраивание выглядит
// как поле без имени.
type container struct {
	base
	str string
}

func main() {

	// При создании структур с помощью литералов нужно
	// явно инициализировать встраивание; здесь встроенный
	// тип служит именем поля.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Мы можем обращаться к полям `base` напрямую через `co`,
	// например, `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Также можно указать полный путь, используя
	// имя встроенного типа.
	fmt.Println("also num:", co.base.num)

	// Поскольку `container` встраивает `base`, методы `base`
	// также становятся методами `container`. Здесь мы вызываем
	// метод, унаследованный от `base`, напрямую через `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Встраивание структур с методами можно использовать
	// для передачи реализации интерфейсов другим структурам.
	// Здесь `container` теперь реализует интерфейс `describer`,
	// потому что встраивает `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
