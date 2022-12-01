package main

//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//// 默认只能有3个主键:只能定义3个参数
//type batchInsert struct {
//	Id   uint   `gorm:"primaryKey"`
//	Name string //`gorm:"primaryKey"`
//
//}
//
//func gorm_batch_insert() {
//	// 连接数据库
//	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	// 返回 error
//	if err != nil {
//		panic(err)
//	}
//	//自动迁移
//	db.AutoMigrate(&batchInsert{})
//	u1 := batchInsert{Id: 1, Name: "张三"}
//
//	var i_data = []batchInsert{{Name:  "张三_1"}, ...., {Name: "张三_10000"}}
//
//	db.Create(&u1) //创建
//	db.CreateInBatches(i_data,100)
//
//	//for _, u1 := range i_data {
//	//	id_value := u1.Name
//	//	fmt.Println("name值:", id_value)
//	//
//	//}
//
//}
