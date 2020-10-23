package code
/**
Mysql支持
 */
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"diabloGin/config"
	"sync"
)
a
//防止并发耗费资源 加入线程锁
var db *gorm.DB
var once sync.Once
func GormInit() *gorm.DB {
	once.Do(func() {
		var err error
		conn, err := gorm.Open("mysql", config.DBUSERNAME+":"+config.DBPASSWORD+"@tcp("+config.DBHOST+")/"+config.DBNAME+"?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println(err)
		}
		conn.LogMode(true)
		db=conn
	})
	return db
	//defer conn.Close()
}

