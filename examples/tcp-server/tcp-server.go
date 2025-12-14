// Пакет `net` предоставляет инструменты для простого
// построения TCP-серверов.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// `net.Listen` запускает сервер на указанной сети
	// (TCP) и адресе (порт 8090 на всех интерфейсах).
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// Закрываем listener для освобождения порта
	// при завершении приложения.
	defer listener.Close()

	// Бесконечный цикл для приёма новых клиентских соединений.
	for {
		// Ожидаем соединение.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// Используем горутину для обработки соединения,
		// чтобы основной цикл мог продолжать принимать
		// новые соединения.
		go handleConnection(conn)
	}
}

// `handleConnection` обрабатывает одно клиентское соединение,
// читая одну строку текста от клиента и возвращая ответ.
func handleConnection(conn net.Conn) {
	// Закрытие соединения освобождает ресурсы после
	// завершения взаимодействия с клиентом.
	defer conn.Close()

	// Используем `bufio.NewReader` для чтения одной строки
	// данных от клиента (завершающейся переводом строки).
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// Создаём и отправляем ответ обратно клиенту,
	// демонстрируя двустороннюю коммуникацию.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
