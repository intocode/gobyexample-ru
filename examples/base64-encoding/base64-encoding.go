// Go предоставляет встроенную поддержку
// [кодирования/декодирования base64](https://en.wikipedia.org/wiki/Base64).

package main

// Этот синтаксис импортирует пакет `encoding/base64` с именем
// `b64` вместо стандартного `base64`. Это сэкономит нам
// немного места ниже.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Вот `string`, который мы будем кодировать/декодировать.
	data := "abc123!?$*&()'-=@~"

	// Go поддерживает как стандартный, так и URL-совместимый
	// base64. Вот как кодировать с помощью стандартного
	// кодировщика. Кодировщик требует `[]byte`, поэтому
	// мы преобразуем `string` в этот тип.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Декодирование может вернуть ошибку, которую можно
	// проверить, если не уверен, что входные данные
	// корректно сформированы.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Это кодирует/декодирует с использованием URL-совместимого
	// формата base64.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
