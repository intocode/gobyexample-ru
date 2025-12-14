# Строка кодируется в немного разные значения стандартным
# и URL base64 кодировщиками (завершающий `+` vs `-`),
# но оба декодируются в исходную строку.
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
