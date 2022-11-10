/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.64.10
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31-0ubuntu0.22.04.1)
 Source Host           : localhost:3306
 Source Schema         : cloud-disk

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31-0ubuntu0.22.04.1)
 File Encoding         : 65001

 Date: 10/11/2022 22:27:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for respository_pool
-- ----------------------------
DROP TABLE IF EXISTS `respository_pool`;
CREATE TABLE `respository_pool` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `hash` varchar(32) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `ext` varchar(30) DEFAULT NULL,
  `size` double DEFAULT NULL,
  `path` varbinary(255) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of respository_pool
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for share_basic
-- ----------------------------
DROP TABLE IF EXISTS `share_basic`;
CREATE TABLE `share_basic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `respository_identitt` varchar(36) DEFAULT NULL,
  `expired_time` int DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of share_basic
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `name` varchar(60) DEFAULT NULL,
  `password` varchar(32) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_respository
-- ----------------------------
DROP TABLE IF EXISTS `user_respository`;
CREATE TABLE `user_respository` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL,
  `ext` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_respository
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
