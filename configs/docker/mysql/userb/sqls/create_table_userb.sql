drop table if exists userb;
CREATE TABLE `userb` (
    `user_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `name` varchar(64) COMMENT '姓名',
    `email` varchar(64) COMMENT '邮箱',
    `phone` varchar(32) COMMENT '手机号',
    `address` varchar(128) COMMENT '地址',
    `company` varchar(128) COMMENT '公司',
    `otherjson` varchar(128) COMMENT '其他json',
    `created_at` TIMESTAMP(3) ,
    `updated_at` TIMESTAMP(3) ,
    `deleted_at` TIMESTAMP(3) ,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;