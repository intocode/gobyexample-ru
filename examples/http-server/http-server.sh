# Запускаем сервер в фоновом режиме.
$ go run http-server.go &

# Обращаемся к маршруту `/hello`.
$ curl localhost:8090/hello
привет
