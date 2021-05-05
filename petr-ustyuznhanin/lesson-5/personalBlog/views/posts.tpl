<!DOCTYPE html>

<html>

<head>
    <title>Personal Blog</title>
</head>

<body>
    {{.Title}}
    <hr>
    {{range .Posts}}
    <div>
        <a href="/post/{{.Id}}">{{.Title}}</a>
        
    </div>
    {{end}}

    <hr>
        <a href="/newpost">new post</a>

</body>

</html>