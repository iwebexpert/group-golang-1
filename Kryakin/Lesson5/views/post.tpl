{{template "UIkit"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BeeGo posts!</title>
</head>
<body>
<h1>{{.Posts.Id}}) {{.Posts.Header}}</h1>
    <hr>
    {{.Posts.Text}}
    <br>
    {{.Posts.Date}}
    <hr>
    <form action="">
        <input type="button" value="Edit" onclick="window.location='/posts/';">
        <input type="button" value="Delete" onclick="window.location='/posts/';">
    </form>
<a href="/posts/">Back to Posts!</a>
</body>
</html>

{{define "UIkit"}}
<!-- UIkit CSS -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/css/uikit.min.css" />

<!-- UIkit JS -->
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit-icons.min.js"></script>
{{end}}