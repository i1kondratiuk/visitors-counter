CREATE DATABASE IF NOT EXISTS `vc`;

USE `vc`;

CREATE TABLE IF NOT EXISTS `user`
(
    `id`       int NOT NULL AUTO_INCREMENT,
    `name`     varchar(256) NOT NULL,
    `username` varchar(256) NOT NULL,
    `password` varchar(256) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`),
    UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `visit_log`
(
    `id`       int NOT NULL AUTO_INCREMENT,
    `username` varchar(256) DEFAULT NULL,
    `counter`  int          DEFAULT NULL,
    `type`     varchar(256) DEFAULT NULL,
    `value`    varchar(256) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;