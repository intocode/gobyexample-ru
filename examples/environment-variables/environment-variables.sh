# Запуск программы показывает, что мы получаем значение
# `FOO`, которое установили в программе, но `BAR` пуст.
$ go run environment-variables.go
FOO: 1
BAR: 

# Список ключей в окружении зависит от конкретной машины.
TERM_PROGRAM
PATH
SHELL
...
FOO

# Если сначала установить `BAR` в окружении,
# запущенная программа получит это значение.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
