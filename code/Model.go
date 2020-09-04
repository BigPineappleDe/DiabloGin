package code
/**
Mysql支持
 */
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"diabloGin/config"
)

//初始化
func DbConnect() *sql.DB{
	connect, err := sql.Open("mysql", config.DBUSERNAME+":"+config.DBPASSWORD+"@tcp("+config.DBHOST+")/"+config.DBNAME)
	if nil != err {
		SetLogger("connect db error: "+err.Error(), "error")
	}
	return connect
}
//表名
func TableName(tableName string) string {
	return config.DBPREFIX + tableName
}
