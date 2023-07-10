CREATE TABLE `user` (
    `uid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_name` varchar(64) COMMENT '姓名',
    `phone` varchar(64) COMMENT '手机号',
    `password` varchar(64) COMMENT '密码',
    `extra` text COMMENT '扩展信息',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `status` int DEFAULT 0 COMMENT '0存在，1删除',
    PRIMARY KEY (`uid`)
) ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_general_ci;