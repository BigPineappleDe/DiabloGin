# 欢迎使用基于Gin搭建的DiabloGin骨架

## 项目设置

```
go mod download
```

## 编译或热编译进行开发

```
原生：go run main.go

第三方：bee等
```

 ## 打包发布

```
go build main.go
```


## 骨架目录

```
diabliGin
	|-cache	//缓存目录
	|-common //公共模块
	|-|-function.go //公共方法文件
	|-config	//配置模块
	|-|-config.go //骨架配置中心
	|-controller	//控制器模块
	|-middleware	//中间件模块
	|-model	//模型
	|-route	//路由
	|-vendor
	|-.gitignore	
	|-go.mod
	|-README.md
```

