ALTER TABLE `user_account`
ADD COLUMN `display_name` varchar(255) NOT NULL,
ADD COLUMN `icon_url` text NOT NULL,
ADD COLUMN `sub` varchar(255) NOT NULL UNIQUE;
