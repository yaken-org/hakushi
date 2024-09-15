ALTER TABLE `post`
    ADD COLUMN `image_id` bigint NOT NULL
    AFTER `user_account_id`;
