-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `ID` varchar(150) NOT NULL,
  `Title` varchar(255) NOT NULL,
  `Text` mediumtext NOT NULL,
  `Tags` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `articles` (`ID`, `Title`, `Text`, `Tags`) VALUES
('1234',	'Title2',	'Text2',	'Go'),
('63ec3645-538b-4e29-b04c-975bc01afad4',	'dasdsadsa',	'asdsadsad',	''),
('794f8198-0740-43aa-adc6-1be8c1c9d37a',	'gfdsfdsf',	'sdfdsfdsf',	'sdfdsfdsf'),
('8d5932a3-de7a-4338-bf65-567a1024690a',	'gergdgergergre',	'ergregregre',	'ergregregre');

-- 2020-08-25 16:13:22