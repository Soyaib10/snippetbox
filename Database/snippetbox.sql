-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Sep 11, 2024 at 12:47 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `snippetbox`
--

-- --------------------------------------------------------

--
-- Table structure for table `snippets`
--

CREATE TABLE `snippets` (
  `id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `content` text NOT NULL,
  `created` datetime NOT NULL,
  `expires` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `snippets`
--

INSERT INTO `snippets` (`id`, `title`, `content`, `created`, `expires`) VALUES
(1, 'First autumn morning', 'First autumn morning\nthe mirror I stare into\nshows my father\'s face.\n', '2024-08-29 15:45:52', '2025-09-05 00:00:00'),
(2, 'An old silent pond', 'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n', '2024-08-29 15:46:32', '2025-08-29 15:46:32'),
(3, 'Over the wintry forest', 'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.', '2024-08-29 15:46:49', '2025-08-29 15:46:49'),
(4, 'O snail', 'O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi', '2024-09-02 16:32:42', '2025-08-08 00:00:00'),
(5, 'Tumi Vule Jeyo Na', 'Prithibir joto sukh\r\njoto valoasha\r\nsob i tomake debo\r\nektai i asha\r\ntumi vule jeyo na amake \r\nami valobashi tomake\r\n\r\n\r\n- an unknown singer', '2024-09-05 19:05:23', '2025-09-05 19:05:23'),
(6, 'Condition Now', 'This is me\r\nI have no idea what to do\r\nI am jobless, tutionless\r\nBut I am alive and Alhamdulillah to Allah.', '2024-09-06 14:17:51', '2024-09-07 14:17:51'),
(7, 'Test Title', 'Test content', '2024-09-06 16:09:17', '2025-09-06 16:09:17'),
(9, 'Golang, The Language', 'Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.', '2024-09-08 17:58:34', '2025-09-08 17:58:34'),
(10, 'Myself', 'A vagabond who wants to be a happy man.', '2024-09-10 16:37:52', '2025-09-10 16:37:52'),
(11, 'Konokcapa song', 'Tomare legeche eto je valo \r\ncaad bujhi i ta janeeeee \r\n', '2024-09-11 07:47:26', '2025-09-11 07:47:26');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `NAME` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `hashed_password` char(60) NOT NULL,
  `created` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `NAME`, `email`, `hashed_password`, `created`) VALUES
(1, 'Soyaib Zihad', 'zihad10@gmail.com', '$2a$12$omTL7TB6QUtkJqkWAUHpNO95hkARWKlrY8Y7/7fEQP8oXrwnS5rda', '2024-09-10 11:14:24');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `snippets`
--
ALTER TABLE `snippets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_snippets_created` (`created`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_uc_email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `snippets`
--
ALTER TABLE `snippets`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
