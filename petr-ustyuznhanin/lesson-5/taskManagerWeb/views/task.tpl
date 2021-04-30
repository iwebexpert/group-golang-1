<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
</head>
<body>
<h1>Задача №{{.Task.Id}}</h1>
<div>
    {{.Task.Text}}
</div>
<hr>
<a href="/tasks">Назад к списку задач</a>
</body>
</html>