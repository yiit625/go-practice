CREATE DATABASE gopractice;
USE gopractice;

DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
                        `id` varchar(22) NOT NULL AUTO_INCREMENT,
                        `name` varchar(100) NOT NULL,
                        `path` varchar(100) NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;