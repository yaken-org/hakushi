CREATE TABLE IF NOT EXISTS `post_item` (
    `id` bigint NOT NULL AUTO_INCREMENT,

    `post_id` bigint NOT NULL,
    `product_id` bigint NOT NULL,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
