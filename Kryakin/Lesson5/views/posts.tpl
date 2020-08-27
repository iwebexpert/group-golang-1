{{template "UIkit"}}

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BeeGo posts!</title>
</head>
<body>
    {{.Title}} 
    <hr>
{{range .Posts}}
<div>   
    <a href="/post/{{.Id}}">{{.Text}}</a>
</div>
{{end}}
<div post-id="NewPost">
<input class="uk-input" type="text" name="Header" placeholder="Post Header"/>
<input class="uk-input" type="text" name="Text" placeholder="Post text"/>
<button class="uk-button" uk-button-default>Сохранить</button>
</div>
</body>

{{define "UIkit"}}
<!-- UIkit CSS -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/css/uikit.min.css" />

<!-- UIkit JS -->
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/uikit@3.5.6/dist/js/uikit-icons.min.js"></script>
{{end}}