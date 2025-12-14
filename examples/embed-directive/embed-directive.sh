# Используй эти команды для запуска примера.
# (Примечание: из-за ограничений go playground этот
# пример можно запустить только на локальной машине.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

