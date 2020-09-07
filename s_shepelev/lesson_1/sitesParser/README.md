# SitesParser  

# Description
Concurrently fetch urls from array of pages and check for subQuery  
If subQuery is on page, this site will return in result

# Usage
go run sitesParser -query=yandex -pages=https://ya.ru -pages=https://golang.org/ -pages=https://geekbrains.ru/  

# Task
Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и  
массив ссылок на страницы, по которым стоит произвести поиск ([]string).  
Результатом работы функции должен быть массив строк со ссылками на страницы,   
на которых обнаружен поисковый запрос. Функция должна искать точное соответствие фразе   
в тексте ответа от сервера по каждой из ссылок.  