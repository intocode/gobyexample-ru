// Go предоставляет встроенную поддержку XML и XML-подобных
// форматов с помощью пакета `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant будет сопоставлен с XML. Как и в примерах с JSON,
// теги полей содержат директивы для кодировщика и декодировщика.
// Здесь мы используем некоторые особые возможности пакета XML:
// имя поля `XMLName` определяет имя XML-элемента, представляющего
// эту структуру; `id,attr` означает, что поле `Id` является
// XML-_атрибутом_, а не вложенным элементом.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Генерируем XML, представляющий наше растение; используем
	// `MarshalIndent` для более читаемого вывода.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Чтобы добавить стандартный XML-заголовок к выводу,
	// добавь его явно.
	fmt.Println(xml.Header + string(out))

	// Используй `Unmarshal` для разбора потока байтов с XML
	// в структуру данных. Если XML некорректен или не может
	// быть сопоставлен с Plant, будет возвращена описательная
	// ошибка.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// Тег поля `parent>child>plant` указывает кодировщику
	// вложить все `plant` под `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
