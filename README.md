# Orm (object relational mapping)

## 1.database-sql

实现支持日志分级的log库、显示日志代码对应的文件名和行号

封装对sqlsite数据库的连接、关闭、语句的执行操作方法

## 2.对象表结构映射

添加了dialect层适配不同的数据库（数据类型的改变、表存在）

添加Schema，利用reflect完成结构体到数据库表结构的映射

添加了表create,drop,table exists的基本操作

## 3.记录新增和查询

Clause层实现将结构体构造成完整sql语句//generator构造独立语句，clause调用generator和拼接

insert代码将对象平铺后插入

find代码将平铺的字段构造成对象

## 4.链式操作和更新删除

添加update、delete、count功能

添加链式调用

## 5.钩子函数

实现CRUD前后前后调用钩子

## 6.支持事务

database/sql自带数据库的事务操作、Begin()开始、Commit()或Rollback()方法终止

将事务方法封装成一个Transaction方法，用回调函数传入要执行的数据库操作语句

## 7.数据库迁移

对象变更时，数据库表中的字段自动迁移

仅支持字段增加和删除

## 结构

geeorm -> dialect -> Schema -> Session -> Clause

geeorm:数据库交互前准备，交互后收尾

dialect:抽象出不同数据库的差异

Schema:对象和表的转化

Session:与数据库表交互、执行sql语句，操作数据库表，实现crud代码

Clause:独立子语句的构建和合并成完整sql语句
