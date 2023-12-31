CREATE TABLE `enroll_white_log_0` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_1` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_2` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_3` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_4` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_5` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_6` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
CREATE TABLE `enroll_white_log_7` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',
      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',
      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',
      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',
      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',
      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',
      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',
      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',
      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',
      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,
      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,
      KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';
