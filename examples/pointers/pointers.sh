# `zeroval` не изменяет `i` в `main`,
# а `zeroptr` изменяет, потому что имеет
# ссылку на адрес памяти этой переменной.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
