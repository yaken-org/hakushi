CREATE TABLE IF NOT EXISTS `user_account` (
    `id` int NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL UNIQUE,
    `display_name` varchar(255) NOT NULL,
    `icon_url` TEXT NOT NULL,

    `sub` VARCHAR(255) NOT NULL UNIQUE,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `auth_user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_account_id` int NOT NULL,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `post` (
    `id` int NOT NULL AUTO_INCREMENT,

    `user_account_id` int NOT NULL,
    `image_id` int NOT NULL,

    `title` varchar(255) NOT NULL,
    `content` text NOT NULL,
    `likes` int NOT NULL DEFAULT 0,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `product` (
    `id` int NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL,
    `link` text NOT NULL,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `annotation` (
    `id` int NOT NULL AUTO_INCREMENT,

    `post_id` int NOT NULL,
    `product_id` int NOT NULL,

    `display_name` varchar(255) NOT NULL,

    `x` int NOT NULL,
    `y` int NOT NULL,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `post_id` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `tag` (
    `id` int NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL UNIQUE,

    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `post_tag` (
    `post_id` int NOT NULL,
    `tag_id` int NOT NULL,

    PRIMARY KEY (`post_id`, `tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
