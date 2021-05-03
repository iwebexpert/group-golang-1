<!DOCTYPE html>
<html lang="en">

<head>
    {{template "script.tpl"}}
    <title>{{.Post.Title}}</title>
</head>

<body>
    <h1>{{.Post.Title}}</h1>
    <div>
        {{.Post.Text}}
    </div>
    <div>
        <form post-id="{{.Post.Id}}">
            <input type="text" name="Title" placeholder="Post Title"/>
            <input type="text" name="Text" placeholder="Post text"/>
            <input type="reset" value="Reset">
            <input type="button" value="Save" onclick="updatePost('{{.Post.Id}}')">
            <input type="button" value="Delete" onclick="deletePost('{{.Post.Id}}')">
            <a href="/">Main page</a>
        </form>

    </div>
    <hr>

</body>

</html>