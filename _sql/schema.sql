CREATE TABLE IF NOT EXISTS `asset` (
  `id` int(18) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `name` varchar(256 NOT NULL,
  `usarname` varchar(256 NOT NULL,
  `password` varchar(256 NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
  UNIQUE KEY `usarname_UNIQUE` (`usarname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci