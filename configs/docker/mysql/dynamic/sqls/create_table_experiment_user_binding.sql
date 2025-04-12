drop table if exists experiment_user_binding;
-- 实验组与用户包绑定关系表
CREATE TABLE `experiment_user_binding` (
    `experiment_group_id` varchar(64) NOT NULL COMMENT '实验组ID',
    `user_package_id` varchar(64) NOT NULL COMMENT '用户包ID',
    `created_at` bigint NOT NULL COMMENT '创建时间 (Unix时间戳)',
    `updated_at` bigint NOT NULL COMMENT '更新时间 (Unix时间戳)',
    PRIMARY KEY (`experiment_group_id`, `user_package_id`),
    INDEX idx_experiment_user_binding_experiment_group_id (`experiment_group_id`),
    INDEX idx_experiment_user_binding_user_package_id (`user_package_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
