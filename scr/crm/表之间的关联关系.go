package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//gorm多对多关系建立与关联模式----------------01
/**

表示外键时要注意，如下例：CompanyID 和 Company的ID 字段类型一定要一样，否则外键会出错

多对一 (Belongs to ) 一对多（has many) ，在数据库中时一样的
*/
// `User` 属于 `Company`，`CompanyID` 是外键

type ModelStudents struct {
	gorm.Model
	Id        int
	Age       int
	City      string
	Fruit     string
	Languages []Language `gorm:"many2many:user_languages;"` //这句必须加上，这样才能关联上  这是第一句代码
}

type Language struct {
	gorm.Model
	Name  string
	Login string
}

//var db *gorm.DB

func gorm_association_relationship_between_tables() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库错误", err)
		panic(err)
	}
	db.AutoMigrate(&ModelStudents{})

	//fmt.Printf("成功:%v\n", db)
	var db1 = db
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功打开数据库%v\n", db1)
	var ms = ModelStudents{Id: 1, Age: 2000, City: "北京", Fruit: "苹果"}
	var sl = Language{Name: "王静谁", Login: "已经登陆"}
	//这是第2句代码
	/**
	调用Association().Append()方法为两个记录(表)创建关联


	*/
	db.Model(&ms).Association("Languages").Append(&sl)
	db.Create(&ms)

	db.Model(&ms).Association("Languages").Find(&sl)
	/**  查询到关联类的字段和数据记录
	name   login
	王静谁 已经登陆
	*/
	fmt.Println("查找所有匹配的关联记录:\n", ms)

	//下面的会报错
	//codes := []string{"zh-CN", "en-US", "ja-JP"}
	//db.Model(&ms).Where("code IN ?", codes).Association("Languages").Find(&sl)
	//
	//db.Model(&ms).Where("code IN ?", codes).Order("code desc").Association("Languages").Find(&sl)

	//暂时没发现有什么用
	//db.Select("Id").Delete(&ms)
	//db.Delete("City")
}
