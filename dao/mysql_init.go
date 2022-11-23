package dao

import (
	"fmt"
	"github.com/20gu00/go_bookManager/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败" + err.Error())
	}
	DB = db
	fmt.Println("数据库初始化成功")

	//自动创建表
	//真正的项目应该是有dba创建好数据表,代码进行ddl并不安全
	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("自动创建表失败", err.Error())
	}

}
