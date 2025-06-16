-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 16 juin 2025 à 19:18
-- Version du serveur : 9.1.0
-- Version de PHP : 8.3.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `project_forum`
--

-- --------------------------------------------------------

--
-- Structure de la table `categories`
--

DROP TABLE IF EXISTS `categories`;
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `categories`
--

INSERT INTO `categories` (`id`, `name`) VALUES
(1, 'leagueoflegends'),
(2, 'politic'),
(3, 'automobile'),
(4, 'gaming');

-- --------------------------------------------------------

--
-- Structure de la table `comments`
--

DROP TABLE IF EXISTS `comments`;
CREATE TABLE IF NOT EXISTS `comments` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `thread_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `content` text NOT NULL,
  `nb_like` int DEFAULT '0',
  `nb_dislike` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `thread_id` (`thread_id`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `comment_likes`
--

DROP TABLE IF EXISTS `comment_likes`;
CREATE TABLE IF NOT EXISTS `comment_likes` (
  `user_id` int NOT NULL,
  `comment_id` int NOT NULL,
  `is_like` tinyint(1) NOT NULL,
  PRIMARY KEY (`user_id`,`comment_id`),
  KEY `comment_id` (`comment_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `threads`
--

DROP TABLE IF EXISTS `threads`;
CREATE TABLE IF NOT EXISTS `threads` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `topic_id` int DEFAULT NULL,
  `author_id` int DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `description` text,
  `content` text NOT NULL,
  `nb_like` int DEFAULT '0',
  `nb_dislike` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `topic_id` (`topic_id`),
  KEY `author_id` (`author_id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `threads`
--

INSERT INTO `threads` (`id`, `topic_id`, `author_id`, `name`, `description`, `content`, `nb_like`, `nb_dislike`, `created_at`) VALUES
(1, 1, 1, 'toto', 'toto ttitit', 'titit', 0, 0, '2025-06-13 14:10:12'),
(2, 0, 0, 'F1', 'F1', 'F1', 0, 0, '2025-06-16 17:46:33'),
(3, 3, 1, 'F1', 'F1', 'F1', 0, 0, '2025-06-16 17:47:07'),
(4, 0, 0, 'F1', 'F1', 'F1', 0, 0, '2025-06-16 17:47:45');

-- --------------------------------------------------------

--
-- Structure de la table `thread_likes`
--

DROP TABLE IF EXISTS `thread_likes`;
CREATE TABLE IF NOT EXISTS `thread_likes` (
  `user_id` int NOT NULL,
  `thread_id` int NOT NULL,
  `is_like` tinyint(1) NOT NULL,
  PRIMARY KEY (`user_id`,`thread_id`),
  KEY `thread_id` (`thread_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `topics`
--

DROP TABLE IF EXISTS `topics`;
CREATE TABLE IF NOT EXISTS `topics` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `followers` int DEFAULT '0',
  `category_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `fk_category` (`category_id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

--
-- Déchargement des données de la table `topics`
--

INSERT INTO `topics` (`id`, `name`, `followers`, `category_id`) VALUES
(1, 'Fer IV', 2, 1),
(2, 'T1', 1, 1),
(3, 'F1', 1, 3),
(4, '24Mans', 0, 3);

-- --------------------------------------------------------

--
-- Structure de la table `topic_subscriptions`
--

DROP TABLE IF EXISTS `topic_subscriptions`;
CREATE TABLE IF NOT EXISTS `topic_subscriptions` (
  `user_id` int NOT NULL,
  `topic_id` int NOT NULL,
  `is_subscribed` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`user_id`,`topic_id`),
  KEY `topic_id` (`topic_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `topic_subscriptions`
--

INSERT INTO `topic_subscriptions` (`user_id`, `topic_id`, `is_subscribed`) VALUES
(0, 1, 1),
(0, 2, 1),
(1, 1, 1),
(1, 3, 1);

-- --------------------------------------------------------

--
-- Structure de la table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password_hash` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'user',
  `banned` tinyint(1) NOT NULL DEFAULT '0',
  `is_connect` tinyint(1) NOT NULL DEFAULT '0',
  `followers` int NOT NULL DEFAULT '0',
  `is_follow` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id`, `username`, `password_hash`, `role`, `banned`, `is_connect`, `followers`, `is_follow`) VALUES
(1, 'toto', 'TOTO', 'member', 0, 0, 0, 0),
(3, 'titi', 'toto', 'member', 0, 0, 0, 0);

-- --------------------------------------------------------

--
-- Structure de la table `user_follows`
--

DROP TABLE IF EXISTS `user_follows`;
CREATE TABLE IF NOT EXISTS `user_follows` (
  `follower_id` int NOT NULL,
  `followed_id` int NOT NULL,
  PRIMARY KEY (`follower_id`,`followed_id`),
  KEY `followed_id` (`followed_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `user_friends`
--

DROP TABLE IF EXISTS `user_friends`;
CREATE TABLE IF NOT EXISTS `user_friends` (
  `user1_id` int NOT NULL,
  `user2_id` int NOT NULL,
  PRIMARY KEY (`user1_id`,`user2_id`),
  KEY `user2_id` (`user2_id`)
) ;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
