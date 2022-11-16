CREATE DATABASE `balance_schema` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
CREATE TABLE `balance` (
  `user_id` varchar(45) NOT NULL,
  `user_balance` double DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `report` (
  `accrual_id` varchar(45) NOT NULL,
  `user_id` varchar(45) DEFAULT NULL,
  `service_id` varchar(45) DEFAULT NULL,
  `purchase_id` varchar(45) DEFAULT NULL,
  `price` double DEFAULT NULL,
  PRIMARY KEY (`accrual_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `reserve` (
  `user_id` varchar(45) NOT NULL,
  `service_id` varchar(45) DEFAULT NULL,
  `purchase_id` varchar(45) DEFAULT NULL,
  `price` double DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
