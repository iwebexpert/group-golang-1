DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  ID int unsigned NOT NULL AUTO_INCREMENT,
  Header varchar(255),
  Text text,
UNIQUE KEY `id` (`id`));

INSERT INTO posts VALUES
(null,'Мой первый пост','Жили-были дед да баба.'),
(null,'Мой второй пост','И была у них Курочка Ряба.'),
(null,'Мой третий пост','Снесла курочка яичко, да не простое — золотое.');
