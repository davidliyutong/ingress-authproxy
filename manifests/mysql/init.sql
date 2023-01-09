-- Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE DATABASE  IF NOT EXISTS `authproxy` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `authproxy`;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `instanceID` varchar(32) DEFAULT NULL,
                        `name` varchar(45) NOT NULL,
                        `status` int(1) DEFAULT 1 COMMENT '1:可用，0:不可用',
                        `nickname` varchar(30) NOT NULL,
                        `password` varchar(255) NOT NULL,
                        `email` varchar(256) NOT NULL,
                        `phone` varchar(20) DEFAULT NULL,
                        `isAdmin` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '1: administrator\\\\n0: non-administrator',
                        `extendShadow` longtext DEFAULT NULL,
                        `loginedAt` timestamp NULL DEFAULT NULL COMMENT 'last login time',
                        `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
                        `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_name` (`name`),
                        UNIQUE KEY `instanceID_UNIQUE` (`instanceID`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `secret`
--

DROP TABLE IF EXISTS `secret`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `secret` (
                          `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                          `instanceID` varchar(32) DEFAULT NULL,
                          `name` varchar(45) NOT NULL,
                          `username` varchar(255) NOT NULL,
                          `secretID` varchar(36) NOT NULL,
                          `secretKey` varchar(255) NOT NULL,
                          `expires` int(64) unsigned NOT NULL DEFAULT 1534308590,
                          `description` varchar(255) NOT NULL,
                          `extendShadow` longtext DEFAULT NULL,
                          `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
                          `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `instanceID_UNIQUE` (`instanceID`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `policy`
--

DROP TABLE IF EXISTS `policy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `policy` (
                          `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                          `instanceID` varchar(32) DEFAULT NULL,
                          `name` varchar(45) NOT NULL,
                          `username` varchar(255) NOT NULL,
                          `policyShadow` longtext DEFAULT NULL,
                          `extendShadow` longtext DEFAULT NULL,
                          `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
                          `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `instanceID_UNIQUE` (`instanceID`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
