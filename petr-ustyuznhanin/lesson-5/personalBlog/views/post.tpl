<!DOCTYPE html>

<html>

<head>
    {{template "script.tpl"}}
    <title>{{.Post.Title}}</title>
</head>

<body>
    <h1>Постецкий №{{.Post.Id}} {{.Post.Title}}</h1>
    <div>
        {{.Post.Text}}
    </div>
    <div>
        <input type="button" value="Save"  onclick="updatePost('{{.Post.Id}}')">
        <input type="reset" value="Reset" >
        <input type="button" value="Delete" onclick="deletePost('{{.Post.Id}}')">
    </div>
    <hr>
    <a href="/">Main page</a>
</body>

</html>