<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
</head>
<body>
{{.Title}}
<hr>
{{range .Tasks}}
<div>
    <a href="/task/{{.Id}}">{{.Text}}</a>
</div>
{{end}}
</body>
</html>