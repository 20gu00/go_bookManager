package dao

import (
	"fmt"
	"github.com/20gu00/go_bookManager/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//声明全局变量
var DB *gorm.DB

func InitMysql() {
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("初始化数据库失败" + err.Error())
	}
	DB = db
	fmt.Println("初始化数据库成功")
	//自动创建表
	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Println("自动创建表失败:", err)
	}
}
