<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Post.Header}}</title>
    {{template "UIkit.tpl"}}
</head>
<body>
    <a href="/" uk-icon="icon: home"></a>
    <h1>{{.Post.ID}}) {{.Post.Header}}!</h1>
    <div class="uk-card uk-card-default uk-card-body">
        <span>
            {{.Post.Text}}
        </span>
    </div>
</body>
</html>