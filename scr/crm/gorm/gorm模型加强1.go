package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	Name     string
	Email    string
	Birthday time.Time
}
type Blog_Model struct {
	ID int `gorm:"primaryKey"`
	//Author  Author `gorm:"embedded"` //对于正常的结构体字段，你也可以通过标签 embedded 将其嵌入 类似与gorm模型加强类的gorm model
	Author  Author `gorm:"embedded;embeddedPrefix:author_xxx_"` //使用标签 embeddedPrefix 来为 db 中的字段名添加前缀,MySQL里字段显示未author_xxx_name和author_xxx_email
	Upvotes int32
}

// 等效于
/*
type Blog_Model struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}
*/
func gorm5() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 返回 error
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&Blog_Model{})
	///var m = Author{Name: "名字:王强", Email: "2322043783@qq.com"}
	//var u1 = Blog_Model{ID: 120, Upvotes: 23, Author: m} //创建记录并更新给出的字段
	//这一句等同于上面2句
	var u1 = Blog_Model{ID: 1200, Upvotes: 23, Author: Author{
		Name:     "邹惩罚",
		Email:    "236763476734@qq.com",
		Birthday: time.Now(),
	}}
	db.Create(&u1) //创建
	//插入多条数据
	var user_data = []Blog_Model{{ID: 1000}, {ID: 1001}, {ID: 1002}}
	db.CreateInBatches(user_data, 2)
	//result.RowsAffected // 返回插入记录的条数
}
