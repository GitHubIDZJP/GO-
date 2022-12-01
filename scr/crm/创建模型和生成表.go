package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**

表示外键时要注意，如下例：CompanyID 和 Company的ID 字段类型一定要一样，否则外键会出错

多对一 (Belongs to ) 一对多（has many) ，在数据库中时一样的
*/
// `User` 属于 `Company`，`CompanyID` 是外键

type BuildTab struct {
	gorm.Model
	Name string
}

func buildTab() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库错误", err)
		panic(err)
	}
	db.AutoMigrate(&BuildTab{})
	db.Create(&BuildTab{})
	//fmt.Printf("成功:%v\n", db)
}
