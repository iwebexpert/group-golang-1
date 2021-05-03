<script>
    async function updatePost(id) {
        console.log('updatePost()');
        let postForm = document.querySelector(`form[post-id="${id}"]`);
        let postTitle = postForm.querySelector('input[name="Title"]').value;
        let postText = postForm.querySelector('input[name="Text"]').value;

        let data = await fetch(`/post/${id}`, {
            method: 'PUT',
            body: JSON.stringify({
                title: postTitle,
                text: postText,
            }),
        });

        let dataPost = await data.json();
        if (dataPost) {
            console.log(dataPost);
            window.location.reload();
        }
        console.log('dataPost');
    }

    async function deletePost(id) {
        console.log('deletePost()');
        fetch(`/post/${id}`, {
            method: 'DELETE',
        }).then(response => {
            window.location.replace('/');
        });
    }
</script>