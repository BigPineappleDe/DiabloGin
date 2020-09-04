package model

/**
框架模型
*/
import (
	"diabloGin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GoOrmInit() {
	db, _ := gorm.Open("mysql", config.DBUSERNAME+":"+config.DBPASSWORD+"@"+config.DBHOST+"/"+config.DBNAME+"?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
}
