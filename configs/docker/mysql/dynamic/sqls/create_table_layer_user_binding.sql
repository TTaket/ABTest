drop table if exists layer_user_binding;
-- 层与用户包绑定关系表
CREATE TABLE `layer_user_binding` (
    `layer_id` varchar(64) NOT NULL COMMENT '层ID',
    `user_package_id` varchar(64) NOT NULL COMMENT '用户包ID',
    `created_at` bigint NOT NULL COMMENT '创建时间 (Unix时间戳)',
    `updated_at` bigint NOT NULL COMMENT '更新时间 (Unix时间戳)',
    PRIMARY KEY (`layer_id`, `user_package_id`),
    INDEX idx_layer_user_binding_layer_id (`layer_id`),
    INDEX idx_layer_user_binding_user_package_id (`user_package_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
