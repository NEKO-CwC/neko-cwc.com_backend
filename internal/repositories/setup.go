package repositories

import (
	"backend/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup(dbName string) *gorm.DB {
	dsnString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true&loc=Local", config.USERNAME, config.PASSWORD, config.IP, config.PORT, dbName)
	DB, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	return DB
}

var UserInfoDB = Setup("userinfo")
var BlogDB = Setup("blog")
var CommentDB = Setup("comment")
