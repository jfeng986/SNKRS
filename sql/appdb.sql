/*========================>database user <===================================*/
CREATE DATABASE user;
USE user;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT 'Username',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT 'Encrypted user password',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT 'Phone number',
  `question` varchar(100) NOT NULL DEFAULT '' COMMENT 'Password recovery question',
  `answer` varchar(100) NOT NULL DEFAULT '' COMMENT 'Password recovery answer',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_phone` (`phone`),
  UNIQUE KEY `uniq_username` (`username`),
  KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='User Table';

CREATE TABLE `user_receive_address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'User ID',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT 'Recipient name',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT 'Phone number',
  `zip_code` varchar(100) NOT NULL DEFAULT '' COMMENT 'Zip code',
  `state` varchar(100) NOT NULL DEFAULT '' COMMENT 'State',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT 'City',
  `address` varchar(128) NOT NULL DEFAULT '' COMMENT 'Address',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'Whether deleted',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Data creation time [do not assign in code]',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Data update time [do not assign in code]',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='User Receive Address Table';


/*========================>database product <===================================*/
CREATE DATABASE product;
USE product;

CREATE TABLE `product` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Product ID',
    `cateid` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Category ID',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'Product name',
    `subtitle` varchar(200) NOT NULL DEFAULT '' COMMENT 'Product subtitle',
    `images` varchar(1024) NOT NULL DEFAULT '' COMMENT 'Image address, separated by commas',
    `detail` varchar(1024) NOT NULL DEFAULT '' COMMENT 'Product details',
    `price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT 'Price',
    `stock` int(11) NOT NULL DEFAULT 0 COMMENT 'Stock quantity',
    `status` int(6) NOT NULL DEFAULT 1 COMMENT 'Product status: 1-On sale, 2-Off sale, 3-Deleted',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    KEY `ix_cateid` (`cateid`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Product Table';

CREATE TABLE `category` (
    `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Category ID',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT 'Category name',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Product Category Table';

CREATE TABLE `product_operation` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'product id',
  `status` int NOT NULL DEFAULT '1' COMMENT '0-off market, 1-on market',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'creation time',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  PRIMARY KEY (`id`),
  KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='product operation table';


/*========================>database orders <===================================*/
CREATE DATABASE `order`;
USE `order`;

CREATE TABLE `orderlist` (
    `id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Order ID',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'User ID',
    `shippingid` bigint(20) NOT NULL DEFAULT 0 COMMENT 'Receiving information table ID',
    `payment` decimal(20,2) DEFAULT NULL DEFAULT 0 COMMENT 'Actual payment amount',
    `paymenttype` tinyint(4) NOT NULL DEFAULT 1 COMMENT 'Payment type: 1-Online payment',
    `status` smallint(6) NOT NULL DEFAULT 10 COMMENT 'Order status: 0-Canceled, 10-Unpaid, 20-Paid, 30-Pending shipment, 40-Pending receipt, 50-Transaction successful, 60-Transaction closed',
    `payment_time` timestamp NOT NULL COMMENT 'Payment time',
    `send_time` timestamp NOT NULL COMMENT 'Shipment time',
    `end_time` timestamp NOT NULL COMMENT 'Transaction completion time',
    `close_time` timestamp NOT NULL COMMENT 'Transaction closure time',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Order Table';

CREATE TABLE `orderitem` (
     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Order Subtable ID',
     `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT 'Order ID',
     `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'User ID',
     `proid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Product ID',
     `proname` varchar(100) NOT NULL DEFAULT '' COMMENT 'Product name',
     `proimage` varchar(500) NOT NULL DEFAULT '' COMMENT 'Product image address',
     `price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT 'Product price',
     `quantity` int(10) NOT NULL DEFAULT 0 COMMENT 'Product quantity',
     `totalprice` decimal(20,2) NOT NULL DEFAULT 0 COMMENT 'Total product price',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
     PRIMARY KEY (`id`),
     KEY `ix_orderid` (`orderid`),
     KEY `ix_userid` (`userid`),
     KEY `ix_proid` (`proid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Order Detail Table';

CREATE TABLE `shipping` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Receiving Information Table ID',
    `orderid` varchar(64) NOT NULL DEFAULT '' COMMENT 'Order ID',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'User ID',
    `receiver_name` varchar(20) NOT NULL DEFAULT '' COMMENT 'Recipient name',
    `receiver_phone` varchar(20) NOT NULL DEFAULT '' COMMENT 'Recipient landline phone',
    `receiver_state` varchar(20) NOT NULL DEFAULT '' COMMENT 'Province',
    `receiver_city` varchar(20) NOT NULL DEFAULT '' COMMENT 'City',
    `receiver_address` varchar(200) NOT NULL DEFAULT '' COMMENT 'Detailed address',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    KEY `ix_orderid` (`orderid`),
    KEY `ix_userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Receiving Information Table';


/*========================>database comment <===================================*/
CREATE DATABASE `comment`;
USE `comment`;

CREATE TABLE `comment`(
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Comment Table ID',
    `targetid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Comment target ID',
    `userid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Comment user ID',
    `content` varchar(255) NOT NULL DEFAULT '' COMMENT 'Comment content',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`id`),
    KEY `ix_targetid` (`targetid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Comment List';

