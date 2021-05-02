Чтобы запустить проект:

1.  создаст контейнер с паролем для рута 1234 и стандратным портом 3306
 `docker run --name gb-personal-blog -e MYSQL_ROOT_PASSWORD=1234 -p 3306:3306 -d mysql`


2. скачать адмайнер для работы с gb-personal-blog и обращаться к ней по псведониму db через внешний порт 8090, а внутрениий 
8080
`docker run --name gb-adminer-blog --link gb-personal-blog:db -p 8090:8080 -d adminer`

3. Подключиться к адмайнеру через localhost:8090/ с именем пользователя root и паролем 1234
4. создать базу данных personalblog 
5. Дать все привелегии, создать нового пользователя taskmanager и в сервере указать %, пароль 1234, так же серваком можно указать localhost и db
6. создать таблтицу posts Id int autoincrement, taxt varchar 255 utf8-gemeral-ci
 
