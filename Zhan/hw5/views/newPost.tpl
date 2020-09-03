<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My blog</title>
    {{template "UIkit.tpl"}}
</head>
<body>
    <a href="/" uk-icon="icon: home"></a>
    <h1>My new post</h1>
    <div class="uk-card uk-card-default uk-card-body">
        <form action="/newpost/new" method="POST">
            <input class="uk-input" type="text" name="Header" placeholder="Post Header"/>
            <input class="uk-input" type="text" name="Text" placeholder="Post text"/>
            <input type="submit" value="Send" class="uk-button uk-button-default">
            <input type="reset" value="Reset" class="uk-button uk-button-default">
        </form>
    </div>
</body>
</html>