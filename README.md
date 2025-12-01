# Go by Example

Содержимое и инструменты разработки для [Go by Example](https://gobyexample.com),
сайта по изучению Go через аннотированные примеры програм.

### Содержимое

Сайт Go By Example создан на основе комментариев с истоников
файлов из `examples` и рендеринга их с помощью `templates`
в статическую директорию `public`. Программы, на основе которых 
реализован этот пайплайн находятся в `tools`, а также указаны в зависимостях в конфиге `go.mod`.

Сбилженная директория `public` может быть отрендерена любым движком, который отдаёт статический контент.
К примеру, официальный сайт использует S3 и CloudFront.

### Сборка

[![test](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml/badge.svg)](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml)

Чтобы собрать сайт у вас должен быть локально установлен Go. Для сборки:

```console
$ tools/build
```

Для бесконечной сборки:

```console
$ tools/build-loop
```

Запустить сайт локально:

```console
$ tools/serve
```

и открыть `http://127.0.0.1:8000/` в вашем браузере.

### Публикация

Для деплоя сайта:

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

### Лицензирование

Эта работа защищена авторским правом Mark McGranaghan и лицензирована по лицензии:
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go Gopher защищен авторским правом [Renée French](https://reneefrench.blogspot.com/) и лицензирвоана на основе
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Переводы

Переводы сообщества сайта Go by Example доступны по:

* [Китайский](https://gobyexample-cn.github.io/) от [gobyexample-cn](https://github.com/gobyexample-cn)
* [Французский](http://le-go-par-l-exemple.keiruaprod.fr) от [keirua](https://github.com/keirua/gobyexample)
* [Итальянский](https://gobyexampleit.andrearaponi.it/) от [andrearaponi](https://github.com/andrearaponi/gobyexample-it)
* [Японский](http://spinute.org/go-by-example) от [spinute](https://github.com/spinute)
* [Корейский](https://mingrammer.com/gobyexample/) от [mingrammer](https://github.com/mingrammer)
* [Русский](https://gobyexample.com.ru/) от [badkaktus](https://github.com/badkaktus)
* [Украинский](https://butuzov.github.io/gobyexample/) от [butuzov](https://github.com/butuzov/gobyexample)
* [Brazilian Portuguese](https://lcslitx.github.io/GoEmExemplos/) от [lcslitx](https://github.com/LCSLITX)
* [Burmese](https://setkyar.github.io/gobyexample) от [Set Kyar Wa Lar](https://github.com/setkyar/gobyexample)

### Благодарности

Спасибо [Jeremy Ashkenas](https://github.com/jashkenas)
за [Docco](http://jashkenas.github.io/docco/), что стало источником вдохновением проекта.

### FAQ

#### Я нашёл проблему в примерах; что мне делать?

We're very happy to fix problem reports and accept contributions! Please submit
[an issue](https://github.com/mmcgrana/gobyexample/issues) or send a Pull Request.
See `CONTRIBUTING.md` for more details.

#### Какая версия нужна для запуска этих примеров?

Given Go's strong [backwards compatibility guarantees](https://go.dev/doc/go1compat),
we expect the vast majority of examples to work on the latest released version of Go
as well as many older releases going back years.

That said, some examples show off new features added in recent releases; therefore,
it's recommended to try running examples with the latest officially released Go version
(see Go's [release history](https://go.dev/doc/devel/release) for details).

#### Я получаю результат при запуске отличный от тех, что используется в примерах. Пример сломан?

Some of the examples demonstrate concurrent code which has a non-deterministic
execution order. It depends on how the Go runtime schedules its goroutines and
may vary by operating system, CPU architecture, or even Go version.

Similarly, examples that iterate over maps may produce items in a different order
from what you're getting on your machine. This is because the order of iteration
over maps in Go is [not specified and is not guaranteed to be the same from one
iteration to the next](https://go.dev/ref/spec#RangeClause).

It doesn't mean anything is wrong with the example. Typically the code in these
examples will be insensitive to the actual order of the output; if the code is
sensitive to the order - that's probably a bug - so feel free to report it.




