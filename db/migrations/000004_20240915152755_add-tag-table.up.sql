CREATE TABLE IF NOT EXISTS `tag` (
    `id` bigint NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL UNIQUE,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `post_tag` (
    `post_id` bigint NOT NULL,
    `tag_id` bigint NOT NULL,

    PRIMARY KEY (`post_id`, `tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
