-- library
CREATE TABLE `bookstore` (
  `id` varchar(36) NOT NULL DEFAULT '' COMMENT '书籍序列号',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '书籍名称',
  `author` varchar(255) DEFAULT '' COMMENT '书籍作者',
  `publish_date` date DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;