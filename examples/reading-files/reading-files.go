// Чтение и запись файлов — базовые задачи, необходимые для
// многих программ на Go. Сначала рассмотрим несколько примеров
// чтения файлов.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Чтение файлов требует проверки большинства вызовов на ошибки.
// Этот помощник упростит проверку ошибок ниже.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Пожалуй, самая базовая задача чтения файлов —
	// загрузить всё содержимое файла в память.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// Часто нужен больший контроль над тем, как и какие
	// части файла читаются. Для этих задач начни с
	// `Open` файла, чтобы получить значение `os.File`.
	f, err := os.Open(path)
	check(err)

	// Читаем несколько байт с начала файла. Позволяем
	// прочитать до 5 байт, но также отмечаем, сколько
	// на самом деле было прочитано.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Можно также переместиться (`Seek`) к известной позиции
	// в файле и читать (`Read`) оттуда.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Другие методы перемещения — относительно текущей
	// позиции курсора,
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// и относительно конца файла.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// Пакет `io` предоставляет некоторые функции, которые
	// могут быть полезны при чтении файлов. Например, чтение
	// вроде показанного выше можно реализовать более надёжно
	// с помощью `ReadAtLeast`.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Встроенной функции перемотки нет, но
	// `Seek(0, io.SeekStart)` делает это.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// Пакет `bufio` реализует буферизованное чтение,
	// которое может быть полезно как для эффективности
	// при множественных мелких чтениях, так и благодаря
	// дополнительным методам чтения.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Закрой файл, когда закончишь (обычно это планируется
	// сразу после `Open` с помощью `defer`).
	f.Close()
}
