CREATE TABLE `shorturl`
(
    `shorturl` varchar(255) NOT NULL COMMENT 'shorten key',
    `url` varchar(255) NOT NULL,
    PRIMARY KEY(`shorturl`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;