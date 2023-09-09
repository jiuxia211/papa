package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func Init() {
	dsn := "root:lcx0821.@tcp(127.0.0.1:3306)/papa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, //缓存查询语句
		//Logger:      logger.Default.LogMode(logger.Info), // 数据库的日志输出
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		}})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	err = DB.AutoMigrate(&News{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("滴滴，数据库连接成功")
}
