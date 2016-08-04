SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(55) DEFAULT NULL,
  `passwd` varchar(32) DEFAULT NULL,
  `email` varchar(55) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('0', '喵了个咪', 'abc123', 'wenzhenxi@vip.qq.com');
INSERT INTO `user` VALUES ('1', '测试账号', '666666', 'test@vip.com');
