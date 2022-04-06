CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `email` varchar(50) NOT NULL DEFAULT '' COMMENT '注册邮箱',
    `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加密盐值',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `number_unique` (`email`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;