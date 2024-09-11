-- MySQL dump 10.13  Distrib 8.3.0, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: gopatungan_db
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `campaign_images`
--

DROP TABLE IF EXISTS `campaign_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `campaign_images` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `campaign_id` int unsigned DEFAULT NULL,
  `filename` varchar(255) DEFAULT NULL,
  `is_primary` tinyint DEFAULT NULL,
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `campaign_id` (`campaign_id`),
  CONSTRAINT `campaign_images_ibfk_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `campaign_images`
--

LOCK TABLES `campaign_images` WRITE;
/*!40000 ALTER TABLE `campaign_images` DISABLE KEYS */;
INSERT INTO `campaign_images` VALUES (1,1,'campaign-images/satu.jpg',0,'2024-08-27 23:32:54','2024-09-10 00:47:38'),(2,1,'campaign-images/dua.jpg',0,'2024-08-27 23:33:12','2024-09-10 00:47:38'),(5,2,'campaign-images/tiga.jpg',1,'2024-08-27 23:38:59','2024-08-28 12:44:33'),(6,2,'campaign-images/empat.jpg',0,'2024-08-27 23:38:59','2024-08-28 12:44:33'),(7,2,'campaign-images/lima.jpg',0,'2024-08-27 23:45:22','2024-08-28 12:44:33'),(9,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:24:43','2024-09-10 00:47:38'),(10,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:24:47','2024-09-10 00:47:38'),(11,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:24:49','2024-09-10 00:47:38'),(12,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:25:00','2024-09-10 00:47:38'),(13,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:29:07','2024-09-10 00:47:38'),(14,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:29:24','2024-09-10 00:47:38'),(15,1,'images/1-WizzMie.jpg',0,'2024-08-30 15:43:18','2024-09-10 00:47:38'),(16,1,'images/1-WizzMie.jpg',0,'2024-09-09 16:29:13','2024-09-10 00:47:38'),(17,1,'assets/images/1-FotoSlesaiSidang.jpg',1,'2024-09-10 00:47:38','2024-09-10 00:47:38');
/*!40000 ALTER TABLE `campaign_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `campaigns`
--

DROP TABLE IF EXISTS `campaigns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `campaigns` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `long_description` text,
  `goal_amount` int unsigned NOT NULL DEFAULT '0',
  `current_amount` int unsigned NOT NULL DEFAULT '0',
  `benefit` text,
  `backer_count` int unsigned NOT NULL DEFAULT '0',
  `slug` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `campaigns_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `campaigns`
--

LOCK TABLES `campaigns` WRITE;
/*!40000 ALTER TABLE `campaigns` DISABLE KEYS */;
INSERT INTO `campaigns` VALUES (1,1,'Campaign 1','Short Description','LONG DESCRIPTIONNNNNNNNNNNNNNNNNNNN',100000000,0,'1. Dapat makan, 2. Dapat Uang, 3. Dapat kekuasaan',0,'campaign-satu','2024-08-27 22:55:17','2024-08-27 23:38:01'),(2,1,'Campaign 2','Short Description','LONG DESCRIPTIONNNNNNNNNNNNNNNNNNNN',100000000,0,'1. Dapat makan, 2. Dapat Uang, 3. Dapat kekuasaan',0,'campaign-dua','2024-08-27 22:55:17','2024-08-27 23:38:01'),(3,2,'Campaign UMKM Simantabtul ketiga Updated aaa','Sebuah Deskripsi Singkat Campaign UMKM Simantabtul Updated aaaaaa','Sebuah Deskripsi Panjang Campaign UMKM Simantabtul Updated',30000000,0,'1. Dapet A, 2. Dapet B, 3. Dapet C,4. Updated',0,'campaign-tiga','2024-08-27 22:55:17','2024-08-29 14:14:20'),(6,2,'Campaign 4','Short Description','LONG DESCRIPTIONNNNNNNNNNNNNNNNNNNN',100000000,0,'1. Dapat makan, 2. Dapat Uang, 3. Dapat kekuasaan',0,'campaign-empat','2024-08-27 22:55:17','2024-08-27 23:38:01'),(7,2,'Campaign 5','Short Description','LONG DESCRIPTIONNNNNNNNNNNNNNNNNNNN',100000000,0,'1. Dapat makan, 2. Dapat Uang, 3. Dapat kekuasaan',0,'campaign-lima','2024-08-27 22:55:17','2024-08-27 23:38:01'),(8,2,'Campaign 6','Short Des','Long Des',100000000,0,'A',0,'campaign-enam','2024-08-27 23:44:22','2024-08-27 23:44:24'),(9,1,'Galang Dana UMKM Simantap','Galang Dana UMKM Simantap','Galang Dana UMKM Simantap',20000000,0,'Galang Dana UMKM Simantap',0,'galang-dana-umkm-simantap-s-uint-1','2024-08-28 20:34:34','2024-08-28 20:34:34'),(10,1,'Campaign UMKM Simantabtul kedua Updated','Sebuah Deskripsi Singkat Campaign UMKM Simantabtul Updated','Sebuah Deskripsi Panjang Campaign UMKM Simantabtul Updated',30000000,0,'1. Dapet A, 2. Dapet B, 3. Dapet C,4. Updated',0,'galang-dana-umkm-simantap-1','2024-08-28 20:35:25','2024-08-29 13:58:56'),(11,13,'Galang Dana UMKM Simantap','Galang Dana UMKM Simantap','Galang Dana UMKM Simantap',20000000,0,'Galang Dana UMKM Simantap',0,'galang-dana-umkm-simantap-13','2024-08-28 20:36:39','2024-08-28 20:36:39'),(12,1,'Campaign UMKM Simantabtul ketiga','Sebuah Deskripsi Singkat Campaign UMKM Simantabtul','Sebuah Deskripsi Panjang Campaign UMKM Simantabtul',20000000,0,'1. Dapet A, 2. Dapet B, 3. Dapet C',0,'campaign-umkm-simantabtul-ketiga-1','2024-08-28 20:55:39','2024-08-28 20:55:39'),(13,13,'Campaign UMKM Simantabtul keempat','Sebuah Deskripsi Singkat Campaign UMKM Simantabtul','Sebuah Deskripsi Panjang Campaign UMKM Simantabtul',20000000,0,'1. Dapet A, 2. Dapet B, 3. Dapet C',0,'campaign-umkm-simantabtul-keempat-13','2024-08-28 20:58:12','2024-08-28 20:58:12');
/*!40000 ALTER TABLE `campaigns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `campaign_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `amount` int unsigned NOT NULL,
  `status` varchar(50) NOT NULL,
  `code` varchar(100) DEFAULT NULL,
  `payment_url` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sericode` (`code`),
  KEY `campaign_id` (`campaign_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,1,1,100000,'paid',NULL,NULL,'2024-09-02 11:41:43','2024-09-02 15:38:37'),(2,2,1,300000,'pending',NULL,NULL,'2024-09-02 11:42:20','2024-09-02 15:41:20'),(3,2,2,350000,'paid',NULL,NULL,'2024-09-03 20:46:25','2024-09-03 20:46:26'),(4,8,13,200000,'pending','',NULL,'2024-09-05 08:41:16','2024-09-05 08:41:16'),(12,7,13,200000,'pending','ORD20240905-8225',NULL,'2024-09-05 10:21:20','2024-09-05 10:21:20'),(13,7,13,200000,'pending','ORD20240905-4433',NULL,'2024-09-05 10:25:15','2024-09-05 10:25:15'),(14,7,13,200000,'pending','ORD20240906-4965','','2024-09-06 16:30:37','2024-09-06 16:30:37'),(15,7,13,200000,'pending','ORD20240906-2450','','2024-09-06 16:35:39','2024-09-06 16:35:39'),(16,7,13,300000,'pending','ORD20240906-9123','','2024-09-06 16:35:49','2024-09-06 16:35:49'),(17,7,13,300000,'pending','ORD20240906-8023','','2024-09-06 16:36:16','2024-09-06 16:36:16'),(18,7,13,300000,'pending','ORD20240906-0323','','2024-09-06 16:36:17','2024-09-06 16:36:17'),(19,7,13,300000,'pending','ORD20240906-6398','','2024-09-06 16:36:26','2024-09-06 16:36:26'),(20,2,2,300000,'pending','ORD20240906-6004','','2024-09-06 16:37:24','2024-09-06 16:37:24'),(21,2,2,300000,'pending','ORD20240906-8107','','2024-09-06 16:38:30','2024-09-06 16:38:30'),(22,2,2,300000,'pending','ORD20240906-8978','','2024-09-06 16:38:31','2024-09-06 16:38:31'),(23,2,2,300000,'pending','ORD20240906-4745','https://app.sandbox.midtrans.com/snap/v4/redirection/d5c7431f-8020-48c6-a231-bbb87e43f04f','2024-09-06 16:50:41','2024-09-06 16:50:41'),(27,1,2,150000,'pending','ORD20240906-1745','','2024-09-06 17:07:39','2024-09-06 17:07:39'),(28,1,2,150000,'pending','ORD20240906-0271','','2024-09-06 17:21:23','2024-09-06 17:21:23'),(29,1,2,150000,'pending','ORD20240906-6891','https://app.sandbox.midtrans.com/snap/v4/redirection/0687b410-85bb-47e5-be64-ec068caf5d5f','2024-09-06 17:26:17','2024-09-06 17:26:17'),(30,1,2,70000,'pending','ORD20240906-7186','https://app.sandbox.midtrans.com/snap/v4/redirection/4c0e718c-2d23-4c70-9ff3-8149f07fe36e','2024-09-06 17:39:47','2024-09-06 17:39:48'),(31,1,2,70000,'pending','ORD20240906-0576','https://app.sandbox.midtrans.com/snap/v4/redirection/6492efb6-159c-47ae-b4f6-ea14f8068802','2024-09-06 18:00:48','2024-09-06 18:00:49'),(32,13,2,525000,'pending','ORD20240908-3402','https://app.sandbox.midtrans.com/snap/v4/redirection/0bd825dc-6331-4a32-833e-11096d2c43f6','2024-09-08 23:36:57','2024-09-08 23:36:57'),(33,13,2,525000,'pending','ORD20240909-8254','https://app.sandbox.midtrans.com/snap/v4/redirection/ad144009-35f6-42fe-962c-3ad352e2a101','2024-09-09 15:53:06','2024-09-09 15:53:07');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `job` varchar(100) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(50) NOT NULL DEFAULT 'user',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'dapuk','developer','dapuk@gmail.com','$2a$04$89B1oK1sLvtuvBNlVXLRBu/5q/PF3nxo2QIqRlIYWpH409Z/JJ.wS','images/1-PasFoto.jpg','user','2024-08-23 14:03:51','2024-09-10 00:47:27'),(2,'adit','mobdev','adit@gmail.com','$2a$04$89B1oK1sLvtuvBNlVXLRBu/5q/PF3nxo2QIqRlIYWpH409Z/JJ.wS','images/2-FotoBangkit.png','user','2024-08-23 14:03:51','2024-08-27 16:11:18'),(3,'arif','golangdev','arif@gmail.com','$2a$04$89B1oK1sLvtuvBNlVXLRBu/5q/PF3nxo2QIqRlIYWpH409Z/JJ.wS','','user','2024-08-23 14:13:56','2024-08-27 16:10:13'),(4,'Test Simpan','','','','','','2024-08-24 13:44:43','2024-08-24 13:44:43'),(5,'Tes Simpan dari Service','Tester','test@example.com','$2a$04$89B1oK1sLvtuvBNlVXLRBu/5q/PF3nxo2QIqRlIYWpH409Z/JJ.wS','','user','2024-08-24 14:04:45','2024-08-24 14:04:45'),(6,'Nama dari Postman','UI Designer','postman@gmail.com','$2a$04$Mb294Odikq7UbxCbbuvpeegLcYw8MMpSC4/haPSAIYfisMgiQwlwa','','user','2024-08-25 23:00:27','2024-08-25 23:00:27'),(8,'Nama','FullstackDev','postman2@gmail.com','$2a$04$XMamVNT5CKYWAU/.KLhxo.XoJ1PwOap7jORg9PQJVPNNZu/jLBRei','','user','2024-08-26 02:27:10','2024-08-26 02:27:10'),(9,'Nama 1','BackEnd','postman3@gmail.com','$2a$04$xn7Mgm92T0BKLVdq5ud6SO7JKF434lhO1BX7NqE8p76hO2W90cUiG','','user','2024-08-26 02:31:06','2024-08-26 02:31:06'),(10,'Nama Formatter','Formatter','postman4@gmail.com','$2a$04$tKb5fNV4Pd6INa5gayXuvugj4tCoZWbbaDUlJZLYI.qqY27VCvIOi','','user','2024-08-26 10:33:30','2024-08-26 10:33:30'),(13,'test jwt register','devhandal','testjwt@gmail.com','$2a$04$zZ4OfYLryZT0.1K48pxdtuTXYENggh97pBbfPRq4VO9sXe/zhosre','','user','2024-08-27 08:36:55','2024-08-27 08:36:55'),(14,'test jwt register','devhandal','afterrefactor@gmail.com','$2a$04$4WWtizgpXz2HP1/m8EkRSeL3Z1qYlRdmBojKgDDBOXQhA16xvvjWO','','user','2024-09-09 15:52:33','2024-09-09 15:52:33');
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

-- Dump completed on 2024-09-11 18:09:56
