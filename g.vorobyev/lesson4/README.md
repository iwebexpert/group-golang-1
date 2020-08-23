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
> describe posts;
+-----------+------------------+------+-----+---------+----------------+
| Field     | Type             | Null | Key | Default | Extra          |
+-----------+------------------+------+-----+---------+----------------+
| post_id   | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| title     | varchar(255)     | YES  |     | NULL    |                |
| post_data | varchar(255)     | YES  |     | NULL    |                |
+-----------+------------------+------+-----+---------+----------------+
3 rows in set (0.00 sec)
```
