ALTER TABLE `user_account`
ADD COLUMN `display_name` varchar(255) NOT NULL AFTER `name`,
ADD COLUMN `icon_url` text NOT NULL AFTER `display_name`,
ADD COLUMN `sub` varchar(255) NOT NULL UNIQUE AFTER `icon_url`;
