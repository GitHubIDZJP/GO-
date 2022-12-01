package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 默认只能有3个主键:只能定义3个参数
type MODEL struct {
	ID        uint           `gorm:"primaryKey"` //主健  --设置了主键就会自动刷新 1234567一直排下去
	CreatedAt int32          `gorm:"primaryKey"` //
	UpdatedAt time.Time      `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 嵌入结构体
type user_model_data struct {
	gorm.Model //model 类似在外面又嵌套一层呢
	Name       string
}

//等效于
/*
type ModelDef struct {
 ID        uint `gorm:"primaryKey"` //主健
 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt gorm.DeletedAt `gorm:"index"
 Name string

}
*/

func gorm4() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 返回 error
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&user_model_data{})
	//插入一条数据
	var u1 = user_model_data{Name: "模型用户"} //创建记录并更新给出的字段

	db.Create(&u1) //创建

	//result.RowsAffected // 返回插入记录的条数
}
