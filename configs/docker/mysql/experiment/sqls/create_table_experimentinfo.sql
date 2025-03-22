DROP TABLE IF EXISTS `experimentinfo`;
CREATE TABLE `experimentinfo` (
    `experiment_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255),
    `description` VARCHAR(255),
    `status` INT,
    `created_at` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` TIMESTAMP(3) NULL,
    PRIMARY KEY (`experiment_id`),
    KEY `idx_experimentinfo_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;