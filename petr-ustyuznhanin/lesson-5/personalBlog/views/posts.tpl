<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
</head>
<body>
{{.Title}}
<hr>
{{range .Posts}}
<div>
    <a href="/post/{{.Id}}">{{.Text}}</a>
</div>
{{end}}
</body>
</html>