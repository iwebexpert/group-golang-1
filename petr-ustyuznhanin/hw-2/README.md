classicServerExample.go - пример простого сервера на го
goroutines.go - наглядный пример использования горутин
pages/ - для проверки работы сервера


1. Используя функцию для поиска из прошлого практического задания, постройте сервер, который будет принимать JSON с поисковым запросом в POST-запросе и возвращать ответ в виде массива строк в JSON.
{
  "search":"фраза для поиска",
  "sites": [
      "первый сайт",
      "второй сайт"
  ]
}

siteParser.go
Для проверки:
POST https://localhost:8080
in body : {"search":"image","sites": ["https://yandex.ru/","https://google.com/"]}

2. Напишите два роута: один будет записывать информацию в Cookie (например, имя), а второй — получать ее и выводить в ответе на запрос.

serverChi.go

Запросы для проверки:
GET https://localhost:8080/one - Unathorized
POST https://localhost:8080/one в Body прописать value: data.
GET https://localhost:8080/one - data