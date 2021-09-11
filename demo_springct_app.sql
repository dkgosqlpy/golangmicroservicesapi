-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 11, 2021 at 05:43 PM
-- Server version: 8.0.26-0ubuntu0.20.04.2
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `demo_springct_app`
--

-- --------------------------------------------------------

--
-- Table structure for table `courses`
--

CREATE TABLE `courses` (
  `id` int NOT NULL,
  `course_name` varchar(65) DEFAULT NULL,
  `course_prof_name` varchar(25) DEFAULT NULL,
  `description` tinytext NOT NULL,
  `status` tinyint NOT NULL DEFAULT '2',
  `created_dt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `courses`
--

INSERT INTO `courses` (`id`, `course_name`, `course_prof_name`, `description`, `status`, `created_dt`) VALUES
(1, 'BCA', 'Satyen Parikh', 'BCA', 2, '2021-08-19 13:49:04'),
(2, 'MCA', 'Satyen Parikh', 'MCA', 2, '2021-08-19 13:49:16'),
(4, 'MCAdddd', 'Satyen Parikh', 'MCAddd', 2, '2021-08-19 13:49:16'),
(6, 'BBA', 'Heral mam', 'BBA Description', 2, '2021-09-08 16:14:18');

-- --------------------------------------------------------

--
-- Table structure for table `map_students_courses`
--

CREATE TABLE `map_students_courses` (
  `msc_id` int NOT NULL,
  `st_id` int NOT NULL,
  `cs_id` int NOT NULL,
  `enrolled_dt` datetime NOT NULL,
  `status` tinyint NOT NULL DEFAULT '2',
  `created_dt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `map_students_courses`
--

INSERT INTO `map_students_courses` (`msc_id`, `st_id`, `cs_id`, `enrolled_dt`, `status`, `created_dt`) VALUES
(1, 1, 2, '2021-08-19 00:24:03', 2, '2021-08-19 00:27:45'),
(2, 1, 1, '2021-08-19 00:24:03', 2, '2021-08-19 00:27:48'),
(5, 2, 1, '2021-08-19 00:24:03', 2, '2021-08-19 00:27:51'),
(6, 2, 2, '2021-08-19 00:24:03', 2, '2021-08-19 00:27:51'),
(11, 11, 4, '2021-08-19 00:24:03', 2, '2021-08-19 00:27:51');

-- --------------------------------------------------------

--
-- Table structure for table `students`
--

CREATE TABLE `students` (
  `st_id` int NOT NULL,
  `st_name` varchar(25) DEFAULT NULL,
  `st_email` varchar(65) DEFAULT NULL,
  `st_phone` varchar(14) DEFAULT NULL,
  `status` tinyint NOT NULL DEFAULT '2',
  `created_dt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `students`
--

INSERT INTO `students` (`st_id`, `st_name`, `st_email`, `st_phone`, `status`, `created_dt`) VALUES
(1, 'dhananjay', 'dhananjay@gmail.com', '9819545584', 2, '2021-08-18 22:45:16'),
(2, 'arunjay', 'arunjay@gmail.com', '9819545586', 2, '2021-08-18 22:45:16'),
(11, 'ajay', 'dhananjayksharma@gmail.com', '9819545584', 2, '2021-08-19 05:36:16'),
(12, 'CIC', 'dhananjayksharma@gmail.com', '9819545584', 2, '2021-08-24 07:40:11'),
(15, 'CIC', 'dhananjayksharma222@gmail.com', '9819545584', 2, '2021-08-24 07:40:27');

-- --------------------------------------------------------

--
-- Stand-in structure for view `view_students`
-- (See below for the actual view)
--
CREATE TABLE `view_students` (
`st_id` int
,`st_name` varchar(25)
,`st_email` varchar(65)
,`st_phone` varchar(14)
,`enrolled_courses` text
);

-- --------------------------------------------------------

--
-- Structure for view `view_students`
--
DROP TABLE IF EXISTS `view_students`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `view_students`  AS  select `st`.`st_id` AS `st_id`,`st`.`st_name` AS `st_name`,`st`.`st_email` AS `st_email`,`st`.`st_phone` AS `st_phone`,group_concat(`cs`.`course_name` separator ',') AS `enrolled_courses` from ((`students` `st` join `map_students_courses` `msc` on((`st`.`st_id` = `msc`.`st_id`))) join `courses` `cs` on((`cs`.`id` = `msc`.`cs_id`))) group by `st`.`st_id` order by `st`.`st_name` ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `courseUnique` (`course_name`);

--
-- Indexes for table `map_students_courses`
--
ALTER TABLE `map_students_courses`
  ADD PRIMARY KEY (`msc_id`) USING BTREE,
  ADD UNIQUE KEY `st_id_cs_id` (`st_id`,`cs_id`),
  ADD KEY `st_id` (`st_id`),
  ADD KEY `cs_id` (`cs_id`);

--
-- Indexes for table `students`
--
ALTER TABLE `students`
  ADD PRIMARY KEY (`st_id`),
  ADD UNIQUE KEY `uniqueNameEmailPhone` (`st_name`,`st_email`,`st_phone`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `courses`
--
ALTER TABLE `courses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `map_students_courses`
--
ALTER TABLE `map_students_courses`
  MODIFY `msc_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `students`
--
ALTER TABLE `students`
  MODIFY `st_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `map_students_courses`
--
ALTER TABLE `map_students_courses`
  ADD CONSTRAINT `map_students_courses_ibfk_1` FOREIGN KEY (`cs_id`) REFERENCES `courses` (`id`) ON DELETE RESTRICT,
  ADD CONSTRAINT `map_students_courses_ibfk_2` FOREIGN KEY (`st_id`) REFERENCES `students` (`st_id`) ON DELETE RESTRICT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
