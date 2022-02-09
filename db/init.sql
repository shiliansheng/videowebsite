-- 创建用户
CREATE USER 'bee'@'%' IDENTIFIED by 'bee';
-- 创建数据库
create database videowebsite default charset utf8 COLLATE utf8_general_ci;
-- 授权
grant all on videowebsite.* to 'bee'@'%';