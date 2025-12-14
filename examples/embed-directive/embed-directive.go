// `//go:embed` — это [директива
// компилятора](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives),
// которая позволяет включать произвольные файлы и папки
// в бинарный файл Go во время сборки. Подробнее о директиве
// embed читай [здесь](https://pkg.go.dev/embed).
package main

// Импортируй пакет `embed`; если не используешь экспортируемые
// идентификаторы из этого пакета, можно сделать пустой импорт
// с помощью `_ "embed"`.
import (
	"embed"
)

// Директивы `embed` принимают пути относительно директории, содержащей
// исходный файл Go. Эта директива встраивает содержимое файла
// в переменную типа `string`, следующую сразу за ней.
//
//go:embed folder/single_file.txt
var fileString string

// Или встроить содержимое файла в `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// Также можно встраивать несколько файлов или даже папки
// с помощью подстановочных знаков. Здесь используется переменная
// типа [embed.FS](https://pkg.go.dev/embed#FS), который реализует
// простую виртуальную файловую систему.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Выводим содержимое `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Получаем некоторые файлы из встроенной папки.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
