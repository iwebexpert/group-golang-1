<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
</head>
<body>
<h1>Постецкий №{{.Post.Id}}</h1>
<div>
    {{.Post.Text}}
</div>
<hr>
<a href="/posts">Назад к списку постов</a>
</body>
</html>