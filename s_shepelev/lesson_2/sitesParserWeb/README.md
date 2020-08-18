# SitesParserWeb 

# Description
Concurrently fetch urls from array of pages and check for subQuery  
If subQuery is on page, this site will return in result in JSON

# Usage
go run sitesParserWeb.go
curl -i localhost:8080 -d '{"search":"yandex","sites": ["https://yandex.ru/","https://google.com/"]}'

# Task
Используя функцию для поиска из прошлого практического задания, постройте сервер,   
который будет принимать JSON с поисковым запросом в POST-запросе и возвращать ответ в виде массива строк в JSON.   
{  
  "search":"фраза для поиска",  
  "sites": [  
      "первый сайт",  
      "второй сайт"  
  ]  
}  