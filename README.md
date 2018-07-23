# RESTful API of [yueshizhixin.github.io](https://yueshizhixin.github.io)

>[English](README.en.md)

### 简述
- 支持检查http header

### 版本
- v0.0.1 注册与登录 

###依赖
- gin github.com/gin-gonic/gin
- mysql github.com/go-sql-driver/mysql
- gorm github.com/jinzhu/gorm
- session github.com/gin-contrib/sessions
- github.com/gin-contrib/cors
- uuid github.com/satori/go.uuid

### 当前代码待优化
- err错误提示统一
- 数据库类型与model类型的详细对应
- http header 内容与cors
- 热更新
- 改用SnowFlake
- i18n


### 后期
- session安全性
- 抢购nsq
- 高并发
