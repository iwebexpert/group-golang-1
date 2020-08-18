домашнее задание 2.

Задание 1
Команды для консоли для проверки работы программы:
curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"mail\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" http://localhost:8080/post
curl -X POST -H 'Content-Type: application/json' -d "{\"search\":\"programming\",\"sites\":[\"https://ya.ru\",\"https://www.google.ru\",\"https://mail.ru\",\"https://www.rambler.ru\",\"https://golang.org\"]}" http://localhost:8080/post

Задание 2
проверяется через барузер
http://localhost:8080/set_cookie
http://localhost:8080/get_cookie
