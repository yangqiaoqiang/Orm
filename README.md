# Orm (object relational mapping)

## 1.database-sql

实现支持日志分级的log库、显示日志代码对应的文件名和行号

封装对sqlsite数据库的连接、关闭、语句的执行操作方法

## 2.对象表结构映射

添加了dialect层适配不同的数据库（数据类型的改变、表存在）

添加Schema，利用reflect完成结构体到数据库表结构的映射

添加了表create,drop,table exists的基本操作

