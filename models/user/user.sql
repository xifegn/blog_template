CREATE TABLE `users`  (
                          `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `name`        varchar(255)        NOT NULL DEFAULT '' COMMENT '用户姓名',
                          `mobile`      varchar(255)        NOT NULL DEFAULT '' COMMENT '用户电话',
                          `password`    varchar(255)        NOT NULL DEFAULT '' COMMENT '用户密码',
                          `collects`    longtext            CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL,
                          `following`   longtext            CHARACTER SET utf8 COLLATE utf8_unicode_ci NULL,
                          `fans`        int(11) unsigned    NOT NULL DEFAULT 0,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `idx_mobile_unique` (`mobile`)
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

