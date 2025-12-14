// Иногда нашим программам на Go нужно порождать
// другие процессы.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Начнём с простой команды, которая не принимает
	// аргументов или ввода и просто выводит что-то в
	// stdout. Помощник `exec.Command` создаёт объект,
	// представляющий этот внешний процесс.
	dateCmd := exec.Command("date")

	// Метод `Output` запускает команду, ждёт её завершения
	// и собирает её стандартный вывод. Если ошибок не было,
	// `dateOut` будет содержать байты с информацией о дате.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` и другие методы `Command` вернут
	// `*exec.Error`, если была проблема с выполнением
	// команды (например, неверный путь), и `*exec.ExitError`,
	// если команда выполнилась, но завершилась с ненулевым
	// кодом возврата.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		var execErr *exec.Error
		var exitErr *exec.ExitError
		switch {
		case errors.As(err, &execErr):
			fmt.Println("failed executing:", err)
		case errors.As(err, &exitErr):
			exitCode := exitErr.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		default:
			panic(err)
		}
	}

	// Далее рассмотрим немного более сложный случай,
	// где мы передаём данные внешнему процессу через его
	// `stdin` и собираем результаты из его `stdout`.
	grepCmd := exec.Command("grep", "hello")

	// Здесь мы явно получаем пайпы ввода/вывода, запускаем
	// процесс, записываем в него некоторый ввод, читаем
	// результирующий вывод и, наконец, ждём завершения
	// процесса.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Мы опустили проверки ошибок в примере выше, но можно
	// использовать обычный паттерн `if err != nil` для всех
	// них. Также мы собираем только результаты `StdoutPipe`,
	// но можно собирать `StderrPipe` точно так же.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Обрати внимание, что при порождении команд нужно
	// предоставить явно разделённый массив команды и
	// аргументов, а не просто одну строку командной строки.
	// Если хочешь выполнить полную команду со строкой,
	// можно использовать опцию `-c` в `bash`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
