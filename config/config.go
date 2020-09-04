package config

/**
配置文件
*/

//核心配置
const (
	PORT            = 8080     //端口
	LOGGING         = true     //日志模块开关
	LOGGINGFILEURL  = "/cache" //日志路径
	LOGGINGDIRLOG   = "log"    //日常日志路径
	LOGGINGDIRERROR = "error"  //错误日志路径
	ISLOGGINGDIROUT = false    //日志存放到项目外
	ISLOGGINGDIR    = ""       //日志存放到项目外的路径 默认项目根目录上级 若需要更改直接填写路径
	TOKENKEY        = ""       //token识别key
)

//Redis配置
const (
	REDISHOST = "127.0.0.1:6379" //redis地址
)

//数据库配置
const (
	DBHOST     = "127.0.0.1:3306" //数据库地址
	DBNAME     = ""               //数据库名
	DBUSERNAME = ""               //数据库账号
	DBPASSWORD = ""               //数据库密码
	DBPREFIX   = ""               //数据表前缀

	//数据表联盟

)
