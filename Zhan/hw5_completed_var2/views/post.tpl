<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Post.Header}}</title>
    {{template "UIkit.tpl"}}
    {{template "script.tpl"}}
</head>
<body>
    <a href="/" uk-icon="icon: home"></a>
    <h1>Post number {{.Post.Id}}</h1>
    <div class="uk-card uk-card-default uk-card-body">
        <form post-id="{{.Post.Id}}">
            <input class="uk-input" type="text" name="Header" value="{{.Post.Header}}"/>
            <input class="uk-input" type="text" name="Text" value="{{.Post.Text}}"/>
            <input type="button" value="Save" class="uk-button uk-button-default" onclick="updatePost('{{.Post.Id}}')">
            <input type="reset" value="Reset" class="uk-button uk-button-default">
            <input type="button" value="Delete" class="uk-button uk-button-default" onclick="deletePost('{{.Post.Id}}')">
            <a href="/" class="uk-button uk-button-default">Back to posts list!</a>
        </form>
    </div>
</body>
</html>