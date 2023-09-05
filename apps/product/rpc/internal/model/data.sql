CREATE DATABASE product;
USE product;

CREATE TABLE `product` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `cateid` smallint(6) UNSIGNED NOT NULL DEFAULT 0,
    `name` varchar(100) NOT NULL DEFAULT '',
    `subtitle` varchar(200) DEFAULT NULL DEFAULT '',
    `detail` text,
    `price` decimal(20,2) NOT NULL DEFAULT 0,
    `stock` int(11) NOT NULL DEFAULT 0,
    `status` int(6) NOT NULL DEFAULT 1 COMMENT '1 on sale,2 off sale,3 delete',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `ix_cateid` (`cateid`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='product table';


CREATE TABLE `category` (
    `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT,
    `parentid` smallint(6) NOT NULL DEFAULT 0,
    `name` varchar(50) NOT NULL DEFAULT '',
    `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '1 normal,2 delete',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='category table';

