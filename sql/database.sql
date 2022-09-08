CREATE DATABASE app_ranking;

CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` varchar(200) NOT NULL DEFAULT '',
  `name` varchar(200) NOT NULL DEFAULT '',
  `publisher_name` varchar(200) NOT NULL DEFAULT '',
  `icon` varchar(255) NOT NULL DEFAULT '',
  `os` varchar(50) NOT NULL DEFAULT '',
  `release_date` varchar(50) NOT NULL DEFAULT '',
  `updated_date` varchar(50) NOT NULL DEFAULT '',
  `rating` float NOT NULL DEFAULT '0',
  `rating_count` int(11) NOT NULL DEFAULT '0',
  `price` float NOT NULL DEFAULT '0',
  `version` varchar(50) NOT NULL DEFAULT '',
  `downloads` int(11) NOT NULL DEFAULT '0',
  `revenue` int(11) NOT NULL DEFAULT '0',
  `support_url` varchar(255) NOT NULL DEFAULT '',
  `website_url` varchar(255) NOT NULL DEFAULT '',
  `privacy_policy_url` varchar(255) NOT NULL DEFAULT '',
  `feature_graphic` varchar(255) NOT NULL DEFAULT '',
  `short_description` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `rank_TH_game_card0` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rank_date` varchar(10) NOT NULL DEFAULT '',
  `rank` int(11) NOT NULL DEFAULT '0',
  `app_id` varchar(200) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=301 DEFAULT CHARSET=utf8;

CREATE TABLE `rank_TH_game_card1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rank_date` varchar(10) NOT NULL DEFAULT '',
  `rank` int(11) NOT NULL DEFAULT '0',
  `app_id` varchar(200) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=301 DEFAULT CHARSET=utf8;

CREATE TABLE `rank_TH_game_card2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rank_date` varchar(10) NOT NULL DEFAULT '',
  `rank` int(11) NOT NULL DEFAULT '0',
  `app_id` varchar(200) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=301 DEFAULT CHARSET=utf8;