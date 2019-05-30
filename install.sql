CREATE TABLE `messages` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `message` text COMMENT '消息体',
  `status` varchar(10) DEFAULT NULL COMMENT 'SUCCESS => 处理成功 ，FAIL =>处理失败',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;