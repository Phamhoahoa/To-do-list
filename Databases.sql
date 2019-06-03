-- MySQL dump 10.13  Distrib 8.0.12, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: todolist
-- ------------------------------------------------------
-- Server version	8.0.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `start` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `due` longtext COLLATE utf8_unicode_ci NOT NULL,
  `note` longtext COLLATE utf8_unicode_ci NOT NULL,
  `id_user` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `f1_idx` (`id_user`),
  CONSTRAINT `f1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task`
--

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
INSERT INTO `task` VALUES (2,'9','9','9','9',6);
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8_unicode_ci NOT NULL,
  `pass` longtext COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (3,'Lan','123456789'),(4,'Lan','$2a$14$sgOBzaiSTSSdG.R5Z/puKOQyzmZeQfvdL5Ok02xplauuHE5mW0wAi'),(5,'Hoa','$2a$14$qiss75zJLX78zu3KcBdIiOqMg8el3uGAoa0ilXuq.k3YLBCxP3vlq'),(6,'huy','$2a$14$HZSLZXnpCCSEgPZbX4JEe.bnn6nRV1SgSZR4wS1/JCwJzplIgXEOS');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `phone` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `token` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2019-02-15 09:31:53','2019-02-15 09:31:53',NULL,'Vạch kẻ đường','145678910','$2a$10$N9aR.QBO2EenfhxQZ7kfIuWQHbCG63JlwhShWvURd6D0V0EK6bvni',NULL),(2,'2019-02-15 09:34:31','2019-02-15 09:34:31',NULL,'Vạch kẻ đường','1456789102','$2a$10$3vgdtahW1B/5ZI30WvUwae/5fL4gMM2sCXZVcxdda2Dx8pXsP1re.',NULL),(3,'2019-02-15 09:37:45','2019-02-15 09:37:45',NULL,'Hoa','14567891','$2a$10$7numzcL9QYhkgMruCWFxoO2UmuNd1zxZ3rOkxAE3phEtM9YACeMaG',NULL),(4,'2019-02-15 09:50:29','2019-02-15 09:50:29',NULL,'Hoa','14567891','$2a$10$04.//99G.OM2xskyo1lmsuREZU1P6OPAHlWcBr9EIxO7sw/Wr5X2W',NULL),(5,'2019-02-15 09:56:44','2019-02-15 09:56:44',NULL,'Hoa','14567891','$2a$10$Adry2kr1G9s7FjQmAMy8KOLtcL/N55YXF68VM0/GcMPgE0k57sHS6',NULL),(6,'2019-02-15 09:58:11','2019-02-15 09:58:11',NULL,'Hoa','14567891','$2a$10$zeddUWn8i1RnDtsYOoX.eexSA1OIYDqrxP6OlY3mkV0jRMAkOmmL2',NULL),(7,'2019-02-15 10:34:04','2019-02-15 10:34:04',NULL,'lan','123456789','$2a$10$pM5CNMaiARXCb4Z3d9CoUeA/M5HKogpKZyMQeesX2x02FEH5z9FpK',NULL),(8,'2019-02-15 11:49:51','2019-02-15 11:49:51',NULL,'deadline','1234567890','$2a$10$iWPJoQbumq1TV181DkJa9uwZwx0YhXlrMZqwJDO6obrA3RmsHhKQ6',''),(9,'2019-02-15 11:50:48','2019-02-15 11:50:48',NULL,'deadline','1234567890','$2a$10$UBasYrRgTY62V/qSIB54I.kOql4hmVKw5EB10.UvvDQL1A6ayRdmS',''),(10,'2019-02-15 12:49:40','2019-02-15 12:49:40',NULL,'Anh','0123456789','$2a$10$lrC4aeYZkL6uG49k0oOsqeDv8snuEC2XEWzwYdbaokhPzuhRic9Zq',''),(11,'2019-02-15 12:50:28','2019-02-15 12:50:28',NULL,'Anh','0123456788','$2a$10$cl/dN8ogXPxuMTXDRo7s.OUtRf1QBE6d2mmEMQRFoQMvkxZsxtRo.','');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-06-03 15:15:54
