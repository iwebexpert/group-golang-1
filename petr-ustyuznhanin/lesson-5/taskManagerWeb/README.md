Чтобы запустить проект:

1.  создаст контейнер с паролем для рута 1234 и стандратным портом 3306
 `docker run --name gb-task-manager -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 -d mysql`


2. скачать адмайнер для работы с gb-task-manager и обращаться к ней по псведониму db через внешний порт 8090, а внутрениий 
8080
`docker run --name gb-adminer --link gb-task-manager:db -p 8090:8080 -d adminer`

3. Подключиться к адмайнеру через localhost:8090/ с именем пользователя root и паролем 1234
4. создать базу данных taskmanager 
5. Дать все привелегии, создать нового пользователя taskmanager и в сервере указать %, пароль 1234, так же серваком указать localhost и db
6. создать таблтицу TaskItems, column name ID varchar 150, Text varchar 255 utf8-general ci , Completed tinyint 1 (boolean тип).
7. создать таблтицу tasks Id int autoincrement, taxt varchar 255 utf8-gemeral-ci
 
