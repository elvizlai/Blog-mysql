SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT;
SET NAMES utf8;
SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO';


-- DROP DATABASE IF EXISTS `blog`;

-- 导出 blog 的数据库结构
CREATE DATABASE IF NOT EXISTS `blog`
  DEFAULT CHARACTER SET utf8;

USE `blog`;

-- 日志表
CREATE TABLE IF NOT EXISTS `logs` (
  `id`    BIGINT(20)    NOT NULL AUTO_INCREMENT,
  `time`  DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `event` VARCHAR(1000) NOT NULL,
  `mid`   VARCHAR(255)  NOT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id`             BIGINT(20)   NOT NULL,
  `nickname`       VARCHAR(255) NOT NULL DEFAULT '',
  `email`          VARCHAR(255) NOT NULL DEFAULT '',
  `password`       VARCHAR(255) NOT NULL DEFAULT '',
  `salt`           VARCHAR(255) NOT NULL DEFAULT '',
  `email_verified` TINYINT(1)   NOT NULL DEFAULT '0',
  `email_token`    VARCHAR(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `email` (`email`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 用户表操作日志trigger
-- 用户增加
DROP TRIGGER IF EXISTS user_add_log;
CREATE TRIGGER user_add_log AFTER INSERT ON user FOR EACH ROW
  BEGIN
    INSERT INTO logs (event, mid) VALUES (CONCAT('CREATE USER:', new.email), new.id);
  END;
-- 用户修改
DROP TRIGGER IF EXISTS user_upd_log;
CREATE TRIGGER user_upd_log AFTER UPDATE ON user FOR EACH ROW
  BEGIN
    INSERT INTO logs (event, mid)
    VALUES (CONCAT('UPDATE USER:', old.nickname, '->', new.nickname, ' ', old.email, '->', new.email), new.id);
  END;
-- 用户删除
DROP TRIGGER IF EXISTS user_del_log;
CREATE TRIGGER user_del_log AFTER DELETE ON user FOR EACH ROW
  BEGIN
    INSERT INTO logs (event, mid) VALUES (CONCAT('DELETE USER:', old.email), old.id);
  END;


-- 用户token表
CREATE TABLE IF NOT EXISTS `user_token` (
  `id`      BIGINT(20)   NOT NULL AUTO_INCREMENT,
  `token`   VARCHAR(255) NOT NULL DEFAULT '',
  `updated` DATETIME     NOT NULL,
  `user_id` BIGINT(20)   NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- token更新trigger + 登录日志
DROP TRIGGER IF EXISTS TokenUpdate;
CREATE TRIGGER TokenUpdate BEFORE UPDATE ON user_token FOR EACH ROW
  BEGIN
    SET new.updated = CURRENT_TIMESTAMP;
    INSERT INTO logs (event, mid) VALUES (CONCAT('UPDATE TOKEN:', new.token), new.id);
  END;

-- 分类表
CREATE TABLE IF NOT EXISTS `category` (
  `id`           BIGINT(20)   NOT NULL AUTO_INCREMENT,
  `name`         VARCHAR(255) NOT NULL DEFAULT '',
  `created`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated`      DATETIME              DEFAULT NULL,
  `user_id`      BIGINT(20)   NOT NULL,
  `last_user_id` BIGINT(20)            DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `category_created` (`created`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 分类表更新时间触发
DROP TRIGGER IF EXISTS CategoryUpdate;
CREATE TRIGGER CategoryUpdate BEFORE UPDATE ON category FOR EACH ROW
  BEGIN
    SET new.updated = CURRENT_TIMESTAMP;
  END;

-- 文章表
CREATE TABLE IF NOT EXISTS `article` (
  `id`           VARCHAR(255)   NOT NULL,
  `category_id`  BIGINT(20)     NOT NULL,
  `title`        VARCHAR(255)   NOT NULL DEFAULT '',
  `tags`         VARCHAR(255)   NOT NULL DEFAULT '',
  `abstract`     VARCHAR(6000)  NOT NULL DEFAULT '',
  `content`      VARCHAR(10000) NOT NULL DEFAULT '',
  `is_draft`     TINYINT(1)     NOT NULL DEFAULT '0',
  `is_del`       TINYINT(1)     NOT NULL DEFAULT '0',
  `created`      DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated`      DATETIME                DEFAULT NULL,
  `views`        BIGINT(20)     NOT NULL DEFAULT '0',
  `user_id`      BIGINT(20)     NOT NULL,
  `last_user_id` BIGINT(20)              DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `article_created` (`created`),
  KEY `article_updated` (`updated`),
  KEY `article_views` (`views`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 添加文章
DROP TRIGGER IF EXISTS article_add_log;
CREATE TRIGGER article_add_log AFTER INSERT ON article FOR EACH ROW
  BEGIN
    INSERT INTO logs (event, mid) VALUES (CONCAT('CREATE ARTICLE:', new.title), new.id);
  END;
-- 修改文章 //todo view数量增加时不应该写log
DROP TRIGGER IF EXISTS article_upd_log;
# CREATE TRIGGER article_upd_log AFTER UPDATE ON article FOR EACH ROW
#   BEGIN
#     INSERT INTO logs (event, mid) VALUES (CONCAT('UPDATE ARTICLE:', new.title), new.id);
#   END;
-- 文章删除
DROP TRIGGER IF EXISTS article_del_log;
CREATE TRIGGER article_del_log AFTER DELETE ON article FOR EACH ROW
  BEGIN
    INSERT INTO logs (event, mid) VALUES (CONCAT('DELETE ARTICLE:', old.title), old.id);
  END;


-- 创建序列并初始化
CREATE TABLE IF NOT EXISTS sequence (
  name          VARCHAR(50) NOT NULL,
  current_value INT         NOT NULL,
  increment     INT         NOT NULL DEFAULT 1,
  PRIMARY KEY (name)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT IGNORE INTO sequence VALUES ('user', 1, 1);


-- 函数
-- 当前值
DROP FUNCTION IF EXISTS currval;
DELIMITER $
CREATE FUNCTION currval(seq_name VARCHAR(50))
  RETURNS INTEGER
CONTAINS SQL
  BEGIN
    DECLARE value INTEGER;
    SET value = 0;
    SELECT current_value
    INTO value
    FROM sequence
    WHERE name = seq_name;
    RETURN value;
  END$
DELIMITER ;

-- 下一个值
DROP FUNCTION IF EXISTS nextval;
DELIMITER $
CREATE FUNCTION nextval(seq_name VARCHAR(50))
  RETURNS INTEGER
CONTAINS SQL
  BEGIN
    UPDATE sequence
    SET current_value = current_value + increment
    WHERE name = seq_name;
    RETURN currval(seq_name);
  END$
DELIMITER ;

-- 设置为某值
DROP FUNCTION IF EXISTS setval;
DELIMITER $
CREATE FUNCTION setval(seq_name VARCHAR(50), value INTEGER)
  RETURNS INTEGER
CONTAINS SQL
  BEGIN
    UPDATE sequence
    SET current_value = value
    WHERE name = seq_name;
    RETURN currval(seq_name);
  END$
DELIMITER ;

-- 附件映射表
CREATE TABLE IF NOT EXISTS `attachment` (
  `id`      BIGINT(20)   NOT NULL AUTO_INCREMENT,
  `path`    VARCHAR(255) NOT NULL DEFAULT '',
  `url`     VARCHAR(255) NOT NULL DEFAULT '',
  `created` DATETIME              DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;