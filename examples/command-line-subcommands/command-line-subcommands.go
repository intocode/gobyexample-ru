// Некоторые инструменты командной строки, такие как `go`
// или `git`, имеют много *подкоманд*, каждая со своим
// набором флагов. Например, `go build` и `go get` — две
// разные подкоманды инструмента `go`. Пакет `flag`
// позволяет легко определять простые подкоманды
// со своими флагами.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Объявляем подкоманду с помощью функции `NewFlagSet`
	// и затем определяем новые флаги, специфичные
	// для этой подкоманды.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Для другой подкоманды можно определить другие
	// поддерживаемые флаги.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Подкоманда ожидается как первый аргумент программы.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Проверяем, какая подкоманда вызвана.
	switch os.Args[1] {

	// Для каждой подкоманды разбираем её собственные флаги
	// и получаем доступ к позиционным аргументам в конце.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
