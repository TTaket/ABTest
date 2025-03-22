drop table if exists `experimentgroup`;
CREATE TABLE `experimentgroup` (
    `group_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255),
    `description` VARCHAR(255),
    `allocation` FLOAT,
    `from_experiment_id` BIGINT UNSIGNED,
    `white_list_user_package_id` BIGINT UNSIGNED,
    `created_at` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` TIMESTAMP(3) NULL,

    PRIMARY KEY (group_id),
    KEY `idx_experimentgroup_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb;