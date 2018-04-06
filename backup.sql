-- MySQL dump 10.14  Distrib 5.5.56-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: soccer
-- ------------------------------------------------------
-- Server version	5.5.56-MariaDB

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
-- Table structure for table `coaches`
--

DROP TABLE IF EXISTS `coaches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `coaches` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `middle_name` varchar(255) DEFAULT NULL,
  `date_of_birth` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `coaches`
--

LOCK TABLES `coaches` WRITE;
/*!40000 ALTER TABLE `coaches` DISABLE KEYS */;
INSERT INTO `coaches` VALUES (1,'MICHAEL','PETERSON','MIKE','1990-01-06 05:00:00'),(2,'JIMMY','SONY','J','1988-03-06 05:00:00'),(3,'NONO','PRINCE','P','1990-01-06 05:00:00'),(4,'CLAUDE','NANY','C','1990-01-06 05:00:00'),(5,'NONO','PRINCE','P','1990-01-06 05:00:00'),(6,'DANIEL','PRINCE','P','1990-01-06 05:00:00'),(7,'DANIEL','JEAN','DJ','1980-09-06 05:00:00');
/*!40000 ALTER TABLE `coaches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `match_schedules`
--

DROP TABLE IF EXISTS `match_schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `match_schedules` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `team_home` int(11) NOT NULL,
  `team_away` int(11) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `team_home` (`team_home`),
  KEY `team_away` (`team_away`),
  CONSTRAINT `match_schedules_ibfk_1` FOREIGN KEY (`team_home`) REFERENCES `teams` (`id`),
  CONSTRAINT `match_schedules_ibfk_2` FOREIGN KEY (`team_away`) REFERENCES `teams` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `match_schedules`
--

LOCK TABLES `match_schedules` WRITE;
/*!40000 ALTER TABLE `match_schedules` DISABLE KEYS */;
INSERT INTO `match_schedules` VALUES (2,1,2,'2017-08-20 05:00:30'),(3,2,1,'2017-08-27 08:00:30'),(4,3,1,'2017-09-15 08:00:30'),(5,1,3,'2017-09-30 08:00:30'),(6,2,3,'2017-10-10 08:00:30'),(7,3,2,'2017-10-25 08:00:30');
/*!40000 ALTER TABLE `match_schedules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `players`
--

DROP TABLE IF EXISTS `players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `players` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `middle_name` varchar(255) DEFAULT NULL,
  `position` varchar(255) DEFAULT NULL,
  `date_of_birth` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `players`
--

LOCK TABLES `players` WRITE;
/*!40000 ALTER TABLE `players` DISABLE KEYS */;
INSERT INTO `players` VALUES (1,'BRIAN','JOHN','BJ','Midfielder','1988-05-08 02:00:00'),(2,'BRIAN','JOHN','BJ','Midfielder','1988-05-08 02:00:00'),(3,'MARK','PHIL','MP','GOALKEEPER','1987-09-08 01:00:00'),(4,'SAMUEL','DON','SD','DEFENDER','1990-05-08 02:00:00'),(5,'CHESNEL','PAUL','CP','FORWARD','1996-05-08 02:00:00'),(6,'DANIEL','PETER','DP','ALL','2018-03-18 22:18:34'),(7,'NIDO','TERE','NT','ALL','2018-03-18 22:19:16'),(8,'NIL','REPET','N','ALL','2018-03-18 22:19:37'),(9,'LIN','PAUL','P','ALL','2018-03-18 22:19:53'),(10,'DEN','LAU','DL','ALL','2018-03-18 22:20:11'),(11,'LEIO','PALU','LP','ALL','2018-03-18 22:20:24'),(12,'BIMA','JONA','B','FORWARD','1988-05-08 02:00:00'),(13,'BAMY','JOA','A','FORWARD','1988-05-08 02:00:00'),(14,'BUMO','GOA','A','FORWARD','1988-05-08 02:00:09'),(15,'CONEL','SONU','CO','FORWARD','1988-05-08 02:00:09'),(16,'CENOL','SUNO','CE','FORWARD','1988-05-08 02:30:09'),(17,'NOLCE','NOSU','NO','FORWARD','1988-05-08 08:30:09'),(18,'CEMO','JONU','CJ','FORWARD','1988-05-08 08:30:09'),(19,'COMU','BOGU','BO','DEFENSE','1988-05-08 08:30:09'),(20,'DUY','JUYN','DJ','DEFENSE','1988-05-08 08:30:09'),(21,'YUD','NYJ','DJ','DEFENSE','1988-05-08 09:30:09'),(22,'UDEY','SHAN','U','FORWARD','1988-05-08 09:30:09'),(23,'FULU','GOCHAN','GO','FORWARD','1988-09-08 09:30:09'),(24,'GUMA','YOGI','Y','FORWARD','1988-05-08 10:30:09'),(25,'MAGU','GIU','G','GOALKEEPER','1988-05-08 10:30:09'),(26,'GUNO','JUAN','J','GOALKEEPER','1988-05-08 10:30:09'),(27,'HUMA','SUMA','SU','DEFENSE','1988-05-08 10:30:01'),(28,'JUMA','LUMO','LU','FORWARD','1988-05-08 10:30:01'),(29,'MAJU','MOMU','MO','FORWARD','1988-05-08 10:30:02'),(30,'PAMU','MULA','P','FORWARD','1988-05-08 10:30:05'),(31,'QUMU','SAMO','QU','FORWARD','1988-05-08 10:30:06'),(32,'DON','SEUL','SE','FORWARD','1988-05-08 10:30:16'),(33,'RUMU','JUMO','J','GOALKEEPER','1988-05-08 10:30:09');
/*!40000 ALTER TABLE `players` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rosters`
--

DROP TABLE IF EXISTS `rosters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rosters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `team_id` int(11) NOT NULL,
  `head_coach` int(11) NOT NULL,
  `assistant_coach` int(11) NOT NULL,
  `player_1` int(11) NOT NULL,
  `player_2` int(11) NOT NULL,
  `player_3` int(11) NOT NULL,
  `player_4` int(11) NOT NULL,
  `player_5` int(11) NOT NULL,
  `player_6` int(11) NOT NULL,
  `player_7` int(11) NOT NULL,
  `player_8` int(11) NOT NULL,
  `player_9` int(11) NOT NULL,
  `player_10` int(11) NOT NULL,
  `player_11` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `team_id` (`team_id`),
  KEY `head_coach` (`head_coach`),
  KEY `assistant_coach` (`assistant_coach`),
  KEY `player_1` (`player_1`),
  KEY `player_2` (`player_2`),
  KEY `player_3` (`player_3`),
  KEY `player_4` (`player_4`),
  KEY `player_5` (`player_5`),
  KEY `player_6` (`player_6`),
  KEY `player_7` (`player_7`),
  KEY `player_8` (`player_8`),
  KEY `player_9` (`player_9`),
  KEY `player_10` (`player_10`),
  KEY `player_11` (`player_11`),
  CONSTRAINT `rosters_ibfk_1` FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`),
  CONSTRAINT `rosters_ibfk_10` FOREIGN KEY (`player_7`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_11` FOREIGN KEY (`player_8`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_12` FOREIGN KEY (`player_9`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_13` FOREIGN KEY (`player_10`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_14` FOREIGN KEY (`player_11`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_2` FOREIGN KEY (`head_coach`) REFERENCES `coaches` (`id`),
  CONSTRAINT `rosters_ibfk_3` FOREIGN KEY (`assistant_coach`) REFERENCES `coaches` (`id`),
  CONSTRAINT `rosters_ibfk_4` FOREIGN KEY (`player_1`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_5` FOREIGN KEY (`player_2`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_6` FOREIGN KEY (`player_3`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_7` FOREIGN KEY (`player_4`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_8` FOREIGN KEY (`player_5`) REFERENCES `players` (`id`),
  CONSTRAINT `rosters_ibfk_9` FOREIGN KEY (`player_6`) REFERENCES `players` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rosters`
--

LOCK TABLES `rosters` WRITE;
/*!40000 ALTER TABLE `rosters` DISABLE KEYS */;
INSERT INTO `rosters` VALUES (2,1,1,2,1,2,3,4,5,6,7,8,9,10,11),(3,2,2,3,12,13,14,15,16,17,18,19,20,21,22),(4,3,3,4,23,24,25,26,27,28,29,30,31,32,33);
/*!40000 ALTER TABLE `rosters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `scores`
--

DROP TABLE IF EXISTS `scores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `scores` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `team_away` int(11) NOT NULL,
  `team_away_score` int(11) NOT NULL,
  `team_home` int(11) NOT NULL,
  `team_home_score` int(11) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `team_away` (`team_away`),
  KEY `team_home` (`team_home`),
  CONSTRAINT `scores_ibfk_1` FOREIGN KEY (`team_away`) REFERENCES `teams` (`id`),
  CONSTRAINT `scores_ibfk_2` FOREIGN KEY (`team_home`) REFERENCES `teams` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scores`
--

LOCK TABLES `scores` WRITE;
/*!40000 ALTER TABLE `scores` DISABLE KEYS */;
INSERT INTO `scores` VALUES (1,1,0,2,3,'2017-08-20 05:00:30'),(2,2,1,1,2,'2017-08-27 08:00:30');
/*!40000 ALTER TABLE `scores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teams`
--

DROP TABLE IF EXISTS `teams`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teams` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teams`
--

LOCK TABLES `teams` WRITE;
/*!40000 ALTER TABLE `teams` DISABLE KEYS */;
INSERT INTO `teams` VALUES (1,'LA Galaxy','California','Los Angeles'),(2,'FC Dallas','Texas','Dallas'),(3,'Red Bulls','New York','New York');
/*!40000 ALTER TABLE `teams` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-04-01 22:03:50
