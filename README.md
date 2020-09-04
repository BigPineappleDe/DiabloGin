# DiabloGin

基于GIN搭建的MVC框架，目的是提供一套轻量易用快速完成业务需求的服务端框架。

## 特性

配置中心

全局异常(基于GIN)

日志系统

ORM(基于gorm)

## 1.配置中心

使用本骨架前需前往config/config.go进行数据库、项目、Redis等配置，具体可到文件中查询具体的默认配置。

## 2.全局异常

本模块基于GIN的中间件，具体可以查看middleware/BaseMiddleware.go文件

## 3.日志系统

日志系统目前仅支持IO文件，

模块的核心源码在code/Logger.go，

使用前需前往config/config.go打开日志模块和配置基本的日志路径，使用时仅需要导入对应的code包code.SetLogger(massage,fileName)即可完成日志的存储。

## 4.ORM

使用第三方gorm框架，具体操作请前往相关文档查阅
