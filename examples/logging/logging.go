// Стандартная библиотека Go предоставляет простые
// инструменты для вывода логов из программ Go: пакет
// [log](https://pkg.go.dev/log) для свободного вывода
// и пакет [log/slog](https://pkg.go.dev/log/slog)
// для структурированного вывода.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// Простой вызов функций вроде `Println` из пакета
	// `log` использует _стандартный_ логгер, который
	// уже предварительно настроен для разумного вывода
	// логов в `os.Stderr`. Дополнительные методы вроде
	// `Fatal*` или `Panic*` завершат программу после
	// логирования.
	log.Println("standard logger")

	// Логгеры можно настраивать с помощью _флагов_ для
	// установки формата вывода. По умолчанию стандартный
	// логгер имеет установленные флаги `log.Ldate` и
	// `log.Ltime`, которые собраны в `log.LstdFlags`.
	// Можно изменить флаги, чтобы выводить время
	// с микросекундной точностью, например.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Также поддерживается вывод имени файла и строки,
	// из которой вызвана функция `log`.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Может быть полезно создать пользовательский логгер
	// и передавать его. При создании нового логгера можно
	// установить _префикс_, чтобы отличать его вывод
	// от других логгеров.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Можно установить префикс на существующих логгерах
	// (включая стандартный) с помощью метода `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Логгеры могут иметь пользовательские цели вывода;
	// подойдёт любой `io.Writer`.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Этот вызов записывает вывод лога в `buf`.
	buflog.Println("hello")

	// Это фактически покажет его в стандартном выводе.
	fmt.Print("from buflog:", buf.String())

	// Пакет `slog` предоставляет _структурированный_
	// вывод логов. Например, логирование в формате JSON
	// делается просто.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// Помимо сообщения, вывод `slog` может содержать
	// произвольное количество пар key=value.
	myslog.Info("hello again", "key", "val", "age", 25)
}
