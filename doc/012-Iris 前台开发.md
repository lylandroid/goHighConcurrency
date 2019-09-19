## 前台核心功能开发
    1. 用户登陆页面开发
    2. 前端商品展示功能开发
    3. 秒杀数据控制开发
    
1. 数据库设计

        /*
         Navicat Premium Data Transfer
        
         Source Server         : mysql
         Source Server Type    : MySQL
         Source Server Version : 50645
         Source Host           : localhost:3306
         Source Schema         : imooc
        
         Target Server Type    : MySQL
         Target Server Version : 50645
         File Encoding         : 65001
        
         Date: 18/09/2019 19:54:46
        */
        
        SET NAMES utf8mb4;
        SET FOREIGN_KEY_CHECKS = 0;
        
        -- ----------------------------
        -- Table structure for user
        -- ----------------------------
        DROP TABLE IF EXISTS `user`;
        CREATE TABLE `user`  (
          `ID` int(11) NOT NULL AUTO_INCREMENT,
          `nickName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
          `userName` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
          `password` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
          PRIMARY KEY (`ID`) USING BTREE
        ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;
        
        SET FOREIGN_KEY_CHECKS = 1;