# Запускаем TCP-сервер в фоновом режиме.
$ go run tcp-server.go &

# Отправляем данные и получаем ответ с помощью netcat.
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT

