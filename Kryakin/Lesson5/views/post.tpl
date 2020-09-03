{{template "UIkit"}}
{{template "JS"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BeeGo posts!</title>
</head>
<body>
    <div post-id="{{.Posts.Id}}" class="uk-card uk-card-default uk-card-body">
<h1>{{.Posts.Id}})<input type="text" class="uk-input" name="Header" value="{{.Posts.Header}}"/> </h1>
    <hr>
    <input type="text" class="uk-input" name="Text" value="{{.Posts.Text}}"/> 
    <br>
    {{.Posts.Date}}
    <hr>
</div>
        <button class="uk-button uk-button-default" onclick="updatePost('{{.Posts.Id}}')">Edit</button>
        <button class="uk-button uk-button-default" onclick="deletePost('{{.Posts.Id}}')">Delete</button>
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

{{define "JS"}}
<script>
    async function updatePost(id){
        console.log('updateTask()');
        let taskForm = document.querySelector(`div[post-id="${id}"]`);
        let postHeader = taskForm.querySelector('input[name="Header"]').value;
        let postText = taskForm.querySelector('input[name="Text"]').value;

        let data = await fetch(`/post/${id}`, {
                    method: 'PUT',
                    body: JSON.stringify({
                        header: postHeader,
                        text: postText,
                    }),
                });

        let dataTask = await data.json();
                if(dataTask){
                    console.log(dataTask);
                    window.location.reload();
                }
    }

    async function deletePost(id){
        console.log('deletePost()');

fetch(`/post/${id}`, {
            method: 'DELETE',
        }).then(response => {
            window.location.replace('/posts/');
        });
}

</script>
{{end}} 