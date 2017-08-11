CREATE TABLE `imhg_bug` (
  `bug_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `api` VARCHAR(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口名称 api/public/UploadImage',
  `title` VARCHAR(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'BUG标题',
  `content` TEXT COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'BUG内容',
  `num` INT(5) DEFAULT '0' COMMENT '当日bug重复数',
  `date` VARCHAR(10) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'bug时间',
  `ctime` INT(10) DEFAULT '0' COMMENT '创建时间',
  `utime` INT(10) DEFAULT '0' COMMENT '修改时间',
  `status` TINYINT(1) DEFAULT '0' COMMENT '状态 0未确认 1已确认',
  PRIMARY KEY (`bug_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci