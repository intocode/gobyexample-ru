$ go build command-line-subcommands.go 

# Сначала вызовем подкоманду foo.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Теперь попробуем bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Но bar не примет флаги foo.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Далее рассмотрим переменные окружения — ещё один
# распространённый способ параметризации программ.
