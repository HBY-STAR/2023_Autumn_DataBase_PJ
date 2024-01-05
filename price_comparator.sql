/*
SQLyog Ultimate v12.08 (32 bit)
MySQL - 8.0.35 : Database - price_comparator
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`price_comparator` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `price_comparator`;

/*Table structure for table `admin` */

DROP TABLE IF EXISTS `admin`;

CREATE TABLE `admin` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `commodity` */

DROP TABLE IF EXISTS `commodity`;

CREATE TABLE `commodity` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `default_name` varchar(64) NOT NULL,
  `produce_at` datetime(3) NOT NULL,
  `produce_address` varchar(64) NOT NULL,
  `category` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_commodity_default_name` (`default_name`),
  KEY `idx_commodity_category` (`category`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `commodity_item` */

DROP TABLE IF EXISTS `commodity_item`;

CREATE TABLE `commodity_item` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `commodity_id` bigint NOT NULL,
  `platform_id` bigint NOT NULL,
  `seller_id` bigint NOT NULL,
  `item_name` varchar(64) NOT NULL,
  `price` float NOT NULL,
  `update_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_item` (`commodity_id`,`platform_id`,`seller_id`),
  KEY `idx_commodity_item_item_name` (`item_name`),
  KEY `idx_commodity_item_update_at` (`update_at` DESC),
  KEY `fk_commodity_item_platform` (`platform_id`),
  KEY `fk_commodity_item_seller` (`seller_id`),
  CONSTRAINT `fk_commodity_item_commodity` FOREIGN KEY (`commodity_id`) REFERENCES `commodity` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_commodity_item_platform` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_commodity_item_seller` FOREIGN KEY (`seller_id`) REFERENCES `seller` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `chk_commodity_item_price` CHECK ((`price` > 0))
) ENGINE=InnoDB AUTO_INCREMENT=2002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `favorite` */

DROP TABLE IF EXISTS `favorite`;

CREATE TABLE `favorite` (
  `user_id` bigint NOT NULL,
  `commodity_item_id` bigint NOT NULL,
  `price_limit` float NOT NULL DEFAULT '0',
  `update_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`,`commodity_item_id`),
  UNIQUE KEY `favorite` (`user_id`,`commodity_item_id`),
  KEY `fk_favorite_commodity_item` (`commodity_item_id`),
  CONSTRAINT `fk_favorite_commodity_item` FOREIGN KEY (`commodity_item_id`) REFERENCES `commodity_item` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_favorite_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `commodity_item_id` bigint DEFAULT NULL,
  `current_price` float NOT NULL,
  `create_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_message_user_id` (`user_id`),
  KEY `fk_message_commodity_item` (`commodity_item_id`),
  CONSTRAINT `fk_message_commodity_item` FOREIGN KEY (`commodity_item_id`) REFERENCES `commodity_item` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_message_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1004 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `platform` */

DROP TABLE IF EXISTS `platform`;

CREATE TABLE `platform` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `url` varchar(64) DEFAULT NULL,
  `country` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `price_change` */

DROP TABLE IF EXISTS `price_change`;

CREATE TABLE `price_change` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `commodity_item_id` bigint NOT NULL,
  `new_price` float NOT NULL,
  `update_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_price_change_commodity_item_id` (`commodity_item_id`),
  KEY `price_change` (`commodity_item_id`,`update_at` DESC),
  CONSTRAINT `fk_price_change_commodity_item` FOREIGN KEY (`commodity_item_id`) REFERENCES `commodity_item` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2503 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `seller` */

DROP TABLE IF EXISTS `seller`;

CREATE TABLE `seller` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `email` varchar(64) DEFAULT NULL,
  `address` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_seller_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=202 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `email` varchar(64) DEFAULT NULL,
  `age` tinyint NOT NULL,
  `gender` tinyint(1) NOT NULL,
  `phone` char(13) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `user_jwt_secret` */

DROP TABLE IF EXISTS `user_jwt_secret`;

CREATE TABLE `user_jwt_secret` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `secret` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=202 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
