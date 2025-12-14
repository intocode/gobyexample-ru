# Попробуй запустить код записи файлов.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Затем проверь содержимое записанных файлов.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Далее рассмотрим применение некоторых идей файлового
# ввода-вывода к потокам `stdin` и `stdout`.
