drop table if exists layer;
CREATE TABLE `layer` (
    `id` varchar(64) NOT NULL COMMENT '层域ID',
    `type` tinyint NOT NULL COMMENT '类型（层或域）',
    `name` varchar(128) NOT NULL COMMENT '名称',
    `description` varchar(256) COMMENT '描述',
    `parent_id` varchar(64) COMMENT '父级ID',
    `ratio` float NOT NULL COMMENT '占有父级比例',
    `created_at` bigint NOT NULL COMMENT '创建时间 (Unix时间戳)',
    `updated_at` bigint NOT NULL COMMENT '更新时间 (Unix时间戳)',
    PRIMARY KEY (`id`),
    INDEX idx_layer_parent_id (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;