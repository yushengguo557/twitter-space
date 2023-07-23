CREATE TABLE `im_twitter_space` (
    `id` CHAR(13) NOT NULL COMMENT '主键 Space ID',
    `creator_id` INT DEFAULT 0 COMMENT '创建者id',
    `participant_count` INT DEFAULT 0 COMMENT '参与人数',
    `title` CHAR(128) DEFAULT '' COMMENT 'Space标题',
    `description` VARCHAR(255) DEFAULT '' COMMENT '描述',
    `lang` ENUM('en', 'zh', 'ja', 'ko') DEFAULT 'en' COMMENT 'Space语言',
    `url` VARCHAR(255) DEFAULT '' COMMENT '链接 跳转推特',
    `status` ENUM('0', '1', '2') DEFAULT '1' COMMENT '状态: 0=end 1=live 2=Scheduled',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `ended_at` datetime DEFAULT NULL COMMENT '结束时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `data_status` BIT(1) NOT NULL DEFAULT 1 COMMENT '数据状态:1=正常,0=删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='推特Space';


CREATE TABLE `im_twitter_user` (
   `id` INT unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id 推特用户id',
   `name` CHAR(16) DEFAULT '' COMMENT '姓名',
   `username` CHAR(16) DEFAULT '' COMMENT '用户名',
   `location` VARCHAR(255) DEFAULT '' COMMENT '位置',
   `description` VARCHAR(255) DEFAULT '' COMMENT '描述',
   `profile_image_url` VARCHAR(255) DEFAULT '' COMMENT '头像链接',
   `space_id` CHAR(13) DEFAULT '' COMMENT '所属Space的ID',
   `created_at` datetime DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
   `data_status` BIT(1) DEFAULT 1 COMMENT '数据状态:1=正常,0=删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='推特User';