// Пакет `filepath` предоставляет функции для разбора
// и построения *путей к файлам* переносимым между
// операционными системами способом; например, `dir/file`
// на Linux против `dir\file` на Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` следует использовать для построения путей
	// переносимым способом. Он принимает любое количество
	// аргументов и строит иерархический путь из них.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Всегда используй `Join` вместо ручной конкатенации
	// `/` или `\`. Помимо обеспечения переносимости, `Join`
	// также нормализует пути, удаляя лишние разделители
	// и переходы между директориями.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` и `Base` можно использовать для разделения пути
	// на директорию и файл. Альтернативно, `Split` вернёт
	// оба значения за один вызов.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Можно проверить, является ли путь абсолютным.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Некоторые имена файлов имеют расширения после точки.
	// Можно отделить расширение от таких имён с помощью `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Чтобы получить имя файла без расширения,
	// используй `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` находит относительный путь между *базой* и
	// *целью*. Возвращает ошибку, если цель не может быть
	// выражена относительно базы.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
