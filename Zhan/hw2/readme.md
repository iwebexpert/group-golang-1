домашнее задание 2.
Команды для консоли для проверки работы программы:

Задание 1
curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" https://postman-echo.com/post
curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" http://localhost:8080/post
