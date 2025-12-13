// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting) —
// важный механизм для контроля использования ресурсов
// и поддержания качества сервиса. Go элегантно
// поддерживает rate limiting с помощью горутин,
// каналов и [тикеров](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Сначала рассмотрим базовый rate limiting. Допустим,
	// мы хотим ограничить обработку входящих запросов.
	// Будем обслуживать эти запросы из одноимённого канала.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Канал `limiter` будет получать значение каждые
	// 200 миллисекунд. Это регулятор в нашей схеме
	// rate limiting.
	limiter := time.Tick(200 * time.Millisecond)

	// Блокируясь на получении из канала `limiter` перед
	// обработкой каждого запроса, мы ограничиваем себя
	// до 1 запроса каждые 200 миллисекунд.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Возможно, мы захотим разрешить короткие всплески
	// запросов в нашей схеме rate limiting, сохраняя
	// при этом общий лимит. Этого можно добиться с помощью
	// буферизации канала-лимитера. Этот `burstyLimiter`
	// позволит всплески до 3 событий.
	burstyLimiter := make(chan time.Time, 3)

	// Заполняем канал для представления разрешённых всплесков.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Каждые 200 миллисекунд мы будем пытаться добавить
	// новое значение в `burstyLimiter`, до его лимита в 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Теперь имитируем ещё 5 входящих запросов. Первые
	// 3 из них воспользуются возможностью всплеска
	// `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
