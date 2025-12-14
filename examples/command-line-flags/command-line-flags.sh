# Для экспериментов с программой флагов командной строки
# лучше сначала скомпилировать её, а затем запустить
# полученный бинарный файл напрямую.
$ go build command-line-flags.go

# Попробуй собранную программу, сначала задав
# значения для всех флагов.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Обрати внимание, что если пропустить флаги, они
# автоматически принимают значения по умолчанию.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Позиционные аргументы в конце можно указать
# после любых флагов.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Обрати внимание, что пакет `flag` требует, чтобы все
# флаги шли перед позиционными аргументами (иначе флаги
# будут интерпретированы как позиционные аргументы).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Используй флаги `-h` или `--help` для получения
# автоматически сгенерированной справки по программе.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Если указать флаг, который не был определён в пакете
# `flag`, программа выведет сообщение об ошибке
# и снова покажет текст справки.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
