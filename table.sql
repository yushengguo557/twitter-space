CREATE TABLE `im_twitter_space` (
    `id` char(13) NOT NULL COMMENT '主键 Space ID',
    `creator_id` char(24) DEFAULT '0' COMMENT '创建者id',
    `participant_count` int(11) DEFAULT '0' COMMENT '参与人数',
    `title` char(128) DEFAULT '' COMMENT 'Space标题',
    `tag` char(16) DEFAULT NULL COMMENT '标签: NFT,WEB3,Game,MetaVerse,DeFi',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
    `lang` enum('en','zh','ja','ko','ru','other','pt','es','tr','de','it','da','ro','pl','sv','fr','nb') NOT NULL DEFAULT 'en' COMMENT 'Space语言',
    `url` varchar(255) DEFAULT '' COMMENT '链接 跳转推特',
    `status` enum('live','ended','scheduled') DEFAULT 'live' COMMENT '状态: ''live'',''ended'',''scheduled''',
    `started_at` datetime DEFAULT NULL COMMENT '创建时间',
    `ended_at` datetime DEFAULT NULL COMMENT '结束时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `data_status` bit(1) NOT NULL DEFAULT b'1' COMMENT '数据状态:1=正常,0=删除',
    `scheduled_start` datetime DEFAULT NULL COMMENT '预定开始时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='推特Space';


CREATE TABLE `im_twitter_user` (
   `id` char(24) NOT NULL COMMENT '主键id 推特用户id',
   `name` char(64) DEFAULT '' COMMENT '姓名',
   `username` char(16) DEFAULT '' COMMENT '用户名',
   `location` varchar(255) DEFAULT '' COMMENT '位置',
   `description` varchar(255) DEFAULT '' COMMENT '描述',
   `profile_image_url` varchar(255) DEFAULT '' COMMENT '头像链接',
   `space_id` char(13) DEFAULT '' COMMENT '所属Space的ID',
   `created_at` datetime DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
   `data_status` bit(1) DEFAULT b'1' COMMENT '数据状态:1=正常,0=删除',
   `url` varchar(255) DEFAULT NULL COMMENT '用户主页链接 ( 推特 )',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='推特User';