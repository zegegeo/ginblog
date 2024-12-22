CREATE TABLE `blogs` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,               -- 博客的唯一 ID
    `title` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL, -- 博客标题
    `content` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,       -- 博客内容
    `user_id` BIGINT NOT NULL,                         -- 博客作者的用户 ID
    `create_time` DATETIME NOT NULL,                   -- 博客创建时间
    `update_time` DATETIME NOT NULL,                   -- 博客更新时间
    PRIMARY KEY (`id`) USING BTREE                     -- 主键
) ENGINE = InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
