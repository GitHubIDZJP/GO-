package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 默认只能有3个主键:只能定义3个参数
type UserModel struct {
	Id       int64  `gorm:"column:username;not null;type:int(4) primary key auto_increment;comment:'用户名'"`
	Password string `gorm:"column:password;type:varchar(30);index:idx_name"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`

	Birthday uint `gorm:"column:Birthday"`

	CreatedAt uint `gorm:"column:CreatedAt"`
	UpdatedAt uint `gorm:"column:UpdatedAt"`
}

func gorm3() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 返回 error
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&UserModel{})
	//t := time.Now().Unix()

	//str := strconv.FormatInt(t, 10)
	u1 := UserModel{Id: 10, Password: "1126", CreateTime: 1994, Birthday: 2022, CreatedAt: 2021, UpdatedAt: 20}
	//u1 := UserModel{Id: }
	//u1 := UserModel{Id: 109, Password: "1126", CreateTime: 2022 - 10 - 20}

	//创建记录并更新给出的字段
	//db.Select("Name", "height", " Gender", "Hobby").Create(&Userinfo{})

	db.Create(&u1) //创建

	//result.RowsAffected // 返回插入记录的条数
}
