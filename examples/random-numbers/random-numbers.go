// Пакет `math/rand/v2` в Go предоставляет генерацию
// [псевдослучайных чисел](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// Например, `rand.IntN` возвращает случайный `int` n,
	// где `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` возвращает `float64` `f`,
	// где `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Это можно использовать для генерации случайных float
	// в других диапазонах, например `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Если нужен известный seed, создай новый `rand.Source`
	// и передай его в конструктор `New`. `NewPCG` создаёт
	// новый источник [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator),
	// который требует seed из двух чисел `uint64`.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
