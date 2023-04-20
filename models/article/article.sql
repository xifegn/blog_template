CREATE TABLE `articles` (
                            `id`          bigint unsigned     NOT NULL AUTO_INCREMENT,
                            `uid`         bigint unsigned     NOT NULL DEFAULT 0 COMMENT '用户ID',
                            `cid`         bigint unsigned     NOT NULL DEFAULT 0 COMMENT '类别ID',
                            `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `title` varchar(50) NOT NULL,
                            `content` text NOT NULL,
                            `head_image` varchar(255) DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;