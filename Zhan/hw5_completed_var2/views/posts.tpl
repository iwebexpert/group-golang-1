<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    {{template "UIkit.tpl"}}
</head>
<body>
    <h1>{{.Title}}</h1>
    <div class="uk-card uk-card-default uk-card-body">
        <span>Последние {{len .Posts}} записей</span>

        <div class="uk-card uk-card-body">
            <ul class="uk-list">
                {{range .Posts}}
                <a style="display:block" href="/post/{{.Id}}">
                <div class="uk-card uk-card-default uk-card-body">
                    <h3>{{.Id}}. {{.Header}}</h3>
                    <span>
                        {{.Text}}
                    </span>
                </div>
                </a>
                {{end}}
            </ul>
        </div>
        <a href="/newpost" class="uk-button uk-button-default">NEW Post</a>

    </div>
</body>
</html>