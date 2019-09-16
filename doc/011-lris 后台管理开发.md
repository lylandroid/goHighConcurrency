## 商品后台管理开发
    1. main models编写
    2. mysql 安装
        1. 安装：docker pull mysql:5.6
        2. 创建目录mysql,用于存放后面的相关东西:
            docker mkdir -p ~/mysql/data ~/mysql/logs ~/mysql/conf
                data目录将映射为mysql容器配置的数据文件存放路径
                logs目录将映射为mysql容器的日志目录
                conf目录里的配置文件将映射为mysql容器的配置文件
        3. 运行容器：docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.6

## 二，form.go-用法说明
1. 在Html模板文件中
<ul>
    <li>使用 “.” 来调用结构体中的熟悉（比如：struct.field1）；</li>
    <li>用索引"[]"的方式获取特殊切片和数组中的值。（比如：struct.array[0]）；</li>
    <li>用键值"[]"的方式来来调用map中的key对应的值（比如：struct.map[es-ES]）；</li>
</ul>

    <code>
        <form method="POST">;
          <input type="text" name="Name" value="Sony">;
          <input type="text" name="Location.Country" value="Japan">;
          <input type="text" name="Location.City" value="Tokyo">;
          <input type="text" name="Products[0].Name" value="Playstation 4">;
          <input type="text" name="Products[0].Type" value="Video games">;
          <input type="text" name="Products[1].Name" value="TV Bravia 32">;
          <input type="text" name="Products[1].Type" value="TVs">;
          <input type="text" name="Founders[0]" value="Masaru Ibuka">;
          <input type="text" name="Founders[0]" value="Akio Morita">;
          <input type="text" name="Employees" value="90000">;
          <input type="text" name="public" value="true">;
          <input type="url" name="website" value="http://www.sony.net">;
          <input type="date" name="foundation" value="1946-05-07">;
          <input type="text" name="Interface.ID" value="12">;
          <input type="text" name="Interface.Name" value="Go Programming Language">;
          <input type="submit">;
        </form>;
    </code>
## 三，创建数据库 & product表
    SET NAMES utf8mb4;
    SET FOREIGN_KEY_CHECKS = 0;
    
    -- ----------------------------
    -- Table structure for product
    -- ----------------------------
    DROP TABLE IF EXISTS `product`;
    CREATE TABLE `product`  (
      `ID` int(11) NOT NULL AUTO_INCREMENT,
      `productName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
      `productNum` int(11) DEFAULT NULL,
      `productImage` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
      `productUrl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
      PRIMARY KEY (`ID`) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

    SET FOREIGN_KEY_CHECKS = 1;
## 四，完成订单相关代码，创建订单表
    SET NAMES utf8mb4;
    SET FOREIGN_KEY_CHECKS = 0;
    
    -- ----------------------------
    -- Table structure for order
    -- ----------------------------
    DROP TABLE IF EXISTS `order`;
    CREATE TABLE `order`  (
      `ID` int(11) NOT NULL AUTO_INCREMENT,
      `userID` int(11) DEFAULT NULL,
      `productID` int(11) DEFAULT NULL,
      `orderStatus` int(11) DEFAULT NULL,
      PRIMARY KEY (`ID`) USING BTREE
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;
    
    SET FOREIGN_KEY_CHECKS = 1;
    
    
#### 4.2, 学习gorm
    
    
    
    