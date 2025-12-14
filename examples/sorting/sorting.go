// Пакет `slices` в Go реализует сортировку для встроенных
// и пользовательских типов. Сначала рассмотрим сортировку
// встроенных типов.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Функции сортировки являются обобщёнными и работают
	// с любым _упорядоченным_ встроенным типом. Список
	// упорядоченных типов см. в [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// Пример сортировки `int`.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// С помощью пакета `slices` можно также проверить,
	// отсортирован ли срез.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
