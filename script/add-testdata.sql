-- INSERT INTO `user_account` (`id`, `name`, `display_name`, `icon_url`, `sub`)
-- VALUES (100, 'testuser', 'テストユーザー', 'https://example.com/icon.png', '1453891135'),
--         (101, 'testuser2', 'テストユーザー2', 'https://example.com/icon2.png', '1453891136'),
--         (102, 'testuser3', 'テストユーザー3', 'https://example.com/icon3.png', '1453891137'),
--         (103, 'testuser4', 'テストユーザー4', 'https://example.com/icon4.png', '1453891138'),
--         (104, 'testuser5', 'テストユーザー5', 'https://example.com/icon5.png', '1453891139'),
--         (105, 'testuser6', 'テストユーザー6', 'https://example.com/icon6.png', '1453891140'),
--         (106, 'testuser7', 'テストユーザー7', 'https://example.com/icon7.png', '1453891141'),
--         (107, 'testuser8', 'テストユーザー8', 'https://example.com/icon8.png', '1453891142'),
--         (108, 'testuser9', 'テストユーザー9', 'https://example.com/icon9.png', '1453891143'),
--         (109, 'testuser10', 'テストユーザー10', 'https://example.com/icon10.png', '1453891144');

INSERT INTO `post` (`id`, `user_account_id`, `image_id`, `title`, `content`, `likes`)
VALUES (100, 100, 100, 'テスト投稿', 'テスト投稿です。', 0),
        (101, 101, 101, 'テスト投稿2', 'テスト投稿です。', 0),
        (102, 102, 102, 'テスト投稿3', 'テスト投稿です。', 0),
        (103, 103, 103, 'テスト投稿4', 'テスト投稿です。', 0),
        (104, 104, 104, 'テスト投稿5', 'テスト投稿です。', 0),
        (105, 105, 105, 'テスト投稿6', 'テスト投稿です。', 0),
        (106, 106, 106, 'テスト投稿7', 'テスト投稿です。', 0),
        (107, 107, 107, 'テスト投稿8', 'テスト投稿です。', 0),
        (108, 108, 108, 'テスト投稿9', 'テスト投稿です。', 0),
        (109, 109, 109, 'テスト投稿10', 'テスト投稿です。', 0),
        (110, 100, 110, 'テスト投稿11', 'テスト投稿です。', 0),
        (111, 101, 111, 'テスト投稿12', 'テスト投稿です。', 0),
        (112, 102, 112, 'テスト投稿13', 'テスト投稿です。', 0),
        (113, 103, 113, 'テスト投稿14', 'テスト投稿です。', 0),
        (114, 104, 114, 'テスト投稿15', 'テスト投稿です。', 0),
        (115, 105, 115, 'テスト投稿16', 'テスト投稿です。', 0),
        (116, 106, 116, 'テスト投稿17', 'テスト投稿です。', 0),
        (117, 107, 117, 'テスト投稿18', 'テスト投稿です。', 0),
        (118, 108, 118, 'テスト投稿19', 'テスト投稿です。', 0),
        (119, 109, 119, 'テスト投稿20', 'テスト投稿です。', 0);

-- INSERT INTO `product` (`id`, `name`, `link`)

