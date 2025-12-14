// В Go есть несколько полезных функций для работы
// с *директориями* в файловой системе.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Создаём новую поддиректорию в текущей рабочей
	// директории.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// При создании временных директорий хорошей практикой
	// является откладывание (`defer`) их удаления. `os.RemoveAll`
	// удалит всё дерево директорий (аналогично `rm -rf`).
	defer os.RemoveAll("subdir")

	// Вспомогательная функция для создания нового пустого файла.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Можно создать иерархию директорий, включая
	// родительские, с помощью `MkdirAll`. Это аналогично
	// команде `mkdir -p`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` выводит содержимое директории, возвращая
	// срез объектов `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` позволяет изменить текущую рабочую директорию,
	// аналогично `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Теперь мы увидим содержимое `subdir/parent/child`
	// при выводе *текущей* директории.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Возвращаемся (`cd`) туда, где начинали.
	err = os.Chdir("../../..")
	check(err)

	// Можно также обойти директорию *рекурсивно*,
	// включая все её поддиректории. `WalkDir` принимает
	// callback-функцию для обработки каждого посещённого
	// файла или директории.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// `visit` вызывается для каждого файла или директории,
// найденных рекурсивно с помощью `filepath.WalkDir`.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
