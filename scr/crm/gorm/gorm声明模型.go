package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID    uint
	Name  string
	Email string
	Age   uint8
	//Blog  struct{}
	Birthday *time.Time
}

type Blog struct {
	ID      int64
	Name    string
	Email   string
	UpVotes int32
	//isBool  bool
	height int32
}

func gorm2() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&User{})
	u1 := User{ID: 10, Name: "王宝强", Email: "zjp15083791732@163.com", Age: 10}
	db.Create(&u1) //创建

	dsn1 := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db1, err1 := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err1 != nil {
		panic(err)
	}
	//自动迁移
	db1.AutoMigrate(&Blog{})
	uc := Blog{ID: 100, Name: "邹总", Email: "2551702817@qq.com", UpVotes: 300, height: 23}
	//u11 := User{ID: 10, Name: "王宝强", Email: "zjp15083791732@163.com", Age: 10}
	db.Create(&uc) //创建

}
