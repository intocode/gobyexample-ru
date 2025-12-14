# Для экспериментов с аргументами командной строки лучше
# сначала собрать бинарный файл с помощью `go build`.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Далее рассмотрим более продвинутую обработку
# командной строки с помощью флагов.
