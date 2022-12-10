package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbConn *gorm.DB

// func GetDbInstance() *gorm.DB {
// 	fmt.Println("1")
// 	if dbConn == nil {
// 		fmt.Println("2")
// 		dbConn = createDbInstance()
// 	}
// 	fmt.Println(dbConn)
// 	return dbConn
// }

// func Init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3308)/gormlab?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	DbConn = db
// }

func Init() {
	d, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DbConn = d
}
