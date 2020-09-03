-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `Id` int unsigned NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Article` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `articles` (`Id`, `Title`, `Article`, `Tags`) VALUES
(9,	'1111111111111111111',	'1111111111111111111111',	'111111111111111111111111111'),
(10,	'ыафыаыфаыфаыфа',	'фыаыфаыфаыфа',	'фыаыфаыфафыаыф'),
(11,	'ааыфавыфаыфаыфа',	'фыаыфаыфаыф',	'фыаыфаыфаыфафы');

-- 2020-08-28 11:58:21