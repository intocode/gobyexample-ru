# Обрати внимание, что хотя срезы и массивы — 
# разные типы, `fmt.Println` отображает их 
# похожим образом.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Подробнее о дизайне и реализации срезов в Go читай
# в [статье](https://go.dev/blog/slices-intro) от команды Go.
