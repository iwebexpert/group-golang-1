DROP DATABASE IF EXISTS my_blog;
CREATE DATABASE my_blog;

USE my_blog

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  id int unsigned NOT NULL AUTO_INCREMENT,
  header varchar(255),
  text text,
UNIQUE KEY `id` (`id`));

INSERT INTO posts VALUES
(null,'Мой первый пост','Жили-были дед да баба.'),
(null,'Мой второй пост','И была у них Курочка Ряба.'),
(null,'Мой третий пост','Снесла курочка яичко, да не простое — золотое.');
