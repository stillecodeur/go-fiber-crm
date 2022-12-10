package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbConn *gorm.DB

func Init() {
	d, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DbConn = d
}
