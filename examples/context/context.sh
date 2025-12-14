# Запускаем сервер в фоновом режиме.
$ go run context.go &

# Имитируем клиентский запрос к `/hello`, нажимая
# Ctrl+C вскоре после начала для сигнала отмены.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
