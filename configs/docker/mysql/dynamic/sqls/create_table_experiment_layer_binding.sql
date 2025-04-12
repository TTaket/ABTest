drop table if exists experiment_layer_binding;
-- 实验与层绑定关系表
CREATE TABLE `experiment_layer_binding` (
    `experiment_id` varchar(64) NOT NULL COMMENT '实验ID',
    `layer_id` varchar(64) NOT NULL COMMENT '层ID',
    `ratio` float NOT NULL COMMENT '占有比例',
    `created_at` bigint NOT NULL COMMENT '创建时间 (Unix时间戳)',
    `updated_at` bigint NOT NULL COMMENT '更新时间 (Unix时间戳)',
    PRIMARY KEY (`experiment_id`, `layer_id`),
    INDEX idx_experiment_layer_binding_experiment_id (`experiment_id`),
    INDEX idx_experiment_layer_binding_layer_id (`layer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;