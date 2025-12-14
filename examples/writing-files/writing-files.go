// Запись файлов в Go следует паттернам, аналогичным тем,
// что мы видели ранее при чтении.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Для начала вот как записать строку (или просто
	// байты) в файл.
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// Для более точной записи открой файл для записи.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// Идиоматично откладывать `Close` сразу
	// после открытия файла.
	defer f.Close()

	// Можно записывать (`Write`) срезы байт, как и ожидается.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Также доступен `WriteString`.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Вызови `Sync` для сброса записей в постоянное хранилище.
	f.Sync()

	// `bufio` предоставляет буферизованные писатели
	// в дополнение к буферизованным читателям, которые мы видели ранее.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Используй `Flush`, чтобы убедиться, что все буферизованные
	// операции были применены к базовому писателю.
	w.Flush()

}
