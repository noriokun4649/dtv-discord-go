
-- +migrate Up
CREATE TABLE IF NOT EXISTS `component_version` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `component` TEXT NOT NULL,
    `version` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `component_version_component_idx` (`component`),
    INDEX `component_version_version_idx` (`version`)
);

-- +migrate Down
DROP TABLE `component_version`;
