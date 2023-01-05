SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `aa_user`;
CREATE TABLE `aa_user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
    `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
    `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'MD5密码',
    `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
    `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态，1正常，2停用',
    `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
    `is_staff` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否员工',
    `is_superuser` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否管理员',
    `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
    `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    KEY `idx_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';