-- borrwo_system
CREATE TABLE `borrow_system` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `book_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '书籍号',
  `user_id` bigint NOT NULL COMMENT '借书人',
  `status` tinyint(1) DEFAULT '0' COMMENT '书籍状态，0-未归还，1-已归还',
  `return_plan_date` timestamp NOT NULL COMMENT '预计还书时间',
  `return_date` int DEFAULT '0' COMMENT '实际还书时间',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_idx` (`user_id`,`book_no`) USING BTREE,
  KEY `book_no_idx` (`book_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;