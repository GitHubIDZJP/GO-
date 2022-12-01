package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"time"
)

// 模型结构
type StudentInsertModel struct {
	//注意：key健必须首字母大写，不然下面赋值了和MySQL里小写字母的字段显示不出来
	//gorm.Model
	Id              int `gorm:"primaryKey"`
	XueLi           string
	Name            string
	Age             int
	Class           string
	Sex             string
	Height          int
	PhysicalFitness string
	School          string
	Distance        string //距离
}

//自定义表名

//func (StudentInsertModel) TableName() String {
//	return "StudentInsertModel"
//}

// 部分字段属性插入
func PartialFieldPropertiesAreInserted() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&StudentInsertModel{})

	//stu := StudentInsertModel{
	//	Id:    2,
	//	Name:  "xin",
	//	Age:   20,
	//	Class: "2班",
	//	Sex:   "男",
	//}
	//xueLi: "本科"
	stu := &StudentInsertModel{Id: 1, XueLi: "本科", Name: "we", Age: 2, Class: "2", Sex: "s", Height: 180, PhysicalFitness: "良好", School: "于都一中 ", Distance: "1.5km"}

	var res = db.Create(&stu)

	res.Select("Id", "Name").Create(stu)
	if res.Error != nil { //判断是否插入数据出错
		fmt.Println(res.Error)
	}
	db.Clauses() //关闭连接
}

//控制台点击运行后控制台不会有任何输出
