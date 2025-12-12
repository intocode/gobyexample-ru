# При запуске этой программы сначала мы видим вывод
# блокирующего вызова, затем вывод двух goroutine.
# Вывод goroutine может чередоваться, поскольку они
# выполняются конкурентно runtime'ом Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Далее мы рассмотрим дополнение к goroutine в
# конкурентных программах Go: каналы.
