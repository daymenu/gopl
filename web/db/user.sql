CREATE TABLE `user` 
(
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `userName` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `status` INT NOT NULL DEFAULT 0 COMMENT '1 可用 2 锁定 3 注销',
    `createdAt` INT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updatedAt` INT NOT NULL DEFAULT 0 COMMENT '修改时间',
    `deletedAt` INT NOT NULL DEFAULT 0 COMMENT '删除时间',
    KEY `userName`(`userName`)
);