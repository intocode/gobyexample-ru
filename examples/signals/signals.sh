# При запуске программа заблокируется в ожидании
# сигнала. Нажав `ctrl-C` (что терминал отображает
# как `^C`), мы отправим сигнал `SIGINT`, и программа
# выведет `interrupt`, а затем завершится.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
