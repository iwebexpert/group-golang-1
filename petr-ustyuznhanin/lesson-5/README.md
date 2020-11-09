# Работа с ORM и фреймоврками

# Tips

- Если делаем обработик, т.е. имеется http.ResponseWriter и http.Request, то лучше отдавать код ответа
типа : `w.WriteHeader(http.StatusNotFound)` а не фатал от которого приложение остановится и не 
будет возможности скорректировать запрос.

# Abstract

Установка фремворка без модуля
`export GO111MODULE=on && go get github.com/beego/bee`

изменения Path , чтобы программа знала где искать утилиты
```export PATH=$PATH:`go env GOPATH`/bin```

создать папку проекта в текущей папке `bee new .`

доставить необходимые пакеты `go get`