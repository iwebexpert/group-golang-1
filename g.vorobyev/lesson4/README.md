1. Создайте модели для своих структур в БД.
2. Создайте методы для получения данных из БД по своим моделям.
3. Адаптируйте роуты, которые обрабатывают запросы на получение всех постов, конкретного поста в блоге и страниц редактирования.

/etc/my.cnf - требуется поддержка utf4mb
```
[client]
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4

[mysqld]
character-set-client-handshake = FALSE
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
```

GORM для линивых - он применят схему сам. Но если что, она такая.

```
DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
  `post_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `post_data` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
