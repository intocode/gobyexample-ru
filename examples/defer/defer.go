// _Defer_ используется для гарантированного выполнения
// вызова функции позже, обычно для целей очистки ресурсов.
// `defer` часто используется там, где в других языках
// применяются `ensure` и `finally`.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Допустим, нам нужно создать файл, записать в него данные,
// а затем закрыть. Вот как это можно сделать с помощью
// `defer`.
func main() {

	// Сразу после получения объекта файла с помощью
	// `createFile` мы откладываем закрытие файла через
	// `closeFile`. Это выполнится в конце охватывающей
	// функции (`main`), после завершения `writeFile`.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Важно проверять ошибки при закрытии файла,
	// даже в отложенной функции.
	if err != nil {
		panic(err)
	}
}
