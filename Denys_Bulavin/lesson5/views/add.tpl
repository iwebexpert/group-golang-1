<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=">
    <title>Document</title>
</head>
<body>
    <div >
        <div>
            <form id="query_form" class="form-horizontal form-well" role="form" action="/add" method="post">
            <div>Заголовок</div>
            <div><input type="text" id="title" name="art-title"></div>
            <div>Текст</div>
            <div><textarea type="textarea" id="article" name="art-article"></textarea></div>
            <div>Тэги</div>
            <div><input type="text" id="tags" name="art-tags"></div>
            <div>
                <button>Сохранить</button>
            </div>
        </form>
        </div>
    </div>
</body>
</html>