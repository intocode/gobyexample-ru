// В предыдущем примере мы использовали явную блокировку с
// помощью [мьютексов](mutexes) для синхронизации доступа
// к общему состоянию из нескольких горутин. Другой вариант —
// использовать встроенные средства синхронизации горутин
// и каналов для достижения того же результата. Такой подход
// на основе каналов соответствует идеям Go о разделении
// памяти через взаимодействие, когда каждый фрагмент данных
// принадлежит ровно одной горутине.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// В этом примере состояние будет принадлежать одной
// горутине. Это гарантирует, что данные никогда не будут
// повреждены при конкурентном доступе. Чтобы прочитать
// или записать это состояние, другие горутины будут
// отправлять сообщения горутине-владельцу и получать
// соответствующие ответы. Структуры `readOp` и `writeOp`
// инкапсулируют эти запросы и способ для горутины-владельца
// отправить ответ.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Как и раньше, будем считать, сколько операций выполнено.
	var readOps uint64
	var writeOps uint64

	// Каналы `reads` и `writes` будут использоваться другими
	// горутинами для отправки запросов на чтение и запись
	// соответственно.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Вот горутина, которая владеет `state` — это словарь,
	// как в предыдущем примере, но теперь он приватный для
	// горутины с состоянием. Эта горутина многократно выполняет
	// select по каналам `reads` и `writes`, отвечая на запросы
	// по мере их поступления. Ответ формируется путём выполнения
	// запрошенной операции и последующей отправки значения
	// в канал ответа `resp` для подтверждения успеха
	// (и желаемого значения в случае `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Здесь запускаются 100 горутин для отправки запросов
	// на чтение горутине-владельцу состояния через канал `reads`.
	// Каждое чтение требует создания `readOp`, отправки его
	// через канал `reads` и получения результата через
	// предоставленный канал `resp`.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Также запускаем 10 записей, используя аналогичный
	// подход.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Дадим горутинам поработать секунду.
	time.Sleep(time.Second)

	// Наконец, фиксируем и выводим счётчики операций.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
