package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 定义模型
type TableUsersModel struct {
	//gorm.Model   //调用gorm框架定义好的模型 ,加了这句会默认产尘3个字段 创建时间created_at  更新时间updated_at 删除时间deleted_at
	Id           int
	Name         string
	Age          int //零值类型
	Birthday     time.Time
	Email        string //`gorm:"type:varchar(100);unique_index"`
	Role         string //`gorm:"size:255"`        //设置字段大小255
	MemberNumber string //`gorm:"unique;not null"` //唯一，不为空
	Num          int    //`gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Address      string //`gorm:"index:addr"`      //给address字段创建名为addr的索引
	IgnoreMe     int    //`grom:"-"`               //忽略本字段

}

// 创建模型和生成表
func create_model_and_generate_table() {
	//连接数据库

	//db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库错误", err)
		panic(err)
	}

	db.AutoMigrate(&TableUsersModel{}) //创建表
	//使用表
	field := TableUsersModel{Id: 1, Name: "曾德", Age: 30, Birthday: time.Now(), Email: "163邮箱", Role: "母亲", MemberNumber: "10", Num: 109, Address: "广州", IgnoreMe: 0}
	//创建表
	res := db.Create(&field)
	//判断是否插入数据出错
	if res.Error != nil {

		fmt.Println(res.Error)
	}

	db.Clauses() //关闭连接
}
