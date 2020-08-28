DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  ID int unsigned NOT NULL AUTO_INCREMENT,
  Header varchar(255) NOT NULL,
  Text varchar(255) NOT NULL,
PRIMARY KEY (`ID`));

INSERT INTO posts VALUES
(null,'My first post','content1'),
(null,'My second post','content2'),
(null,'My third post','content3');
