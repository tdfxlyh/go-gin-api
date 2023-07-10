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


CREATE TABLE `friend_relation` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` bigint(20) NOT NULL COMMENT '自己用户id',
    `other_user_id` bigint(20) NOT NULL COMMENT '朋友用户id',
    `verify_message` text COMMENT '验证消息',
    `notes` varchar(64) COMMENT '备注',
    `rela_status` bigint(20) COMMENT '关系状态-1:请求中,2:好友',
    `extra` text COMMENT '扩展信息',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `status` int DEFAULT 0 COMMENT '0存在，1删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE `message_single` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
   `sender_user_id` bigint(20) NOT NULL COMMENT '消息发送方',
   `receiver_user_id` bigint(20) NOT NULL COMMENT '消息接收方',
   `message_type` bigint(20) COMMENT '消息类型 1:文本 2:图片 3:音频 4:视频 5:文件',
   `content` text COMMENT '消息内容',
   `extra` text COMMENT '扩展信息',
   `read_status_info` int DEFAULT 0 COMMENT '判断接收者 0:未读 1:已读',
   `sender_status_info` int DEFAULT 0 COMMENT '发送方信息状态 0:正常 1:删除',
   `receiver_status_info` int DEFAULT 0 COMMENT '接收方信息状态 0:正常 1:删除',
   `message_status_info` int DEFAULT 0 COMMENT '是否撤回 0:正常 1:撤回',
   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
   `status` int DEFAULT 0 COMMENT '0存在，1删除',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_general_ci;