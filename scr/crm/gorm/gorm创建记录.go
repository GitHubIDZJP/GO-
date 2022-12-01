package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/json"
)

// `gorm:"column:username;not null;type:int(4) primary key auto_increment;comment:'用户名'"`
type BaseModel struct {
	Id int
	//Color string //特性：骚
}

// index索引，type指定字段类型，default默认值，comment注释:
type MARKMODEL struct {
	BaseModel
	OpenID string `gorm:"index:idx_mobile;unique;type:varchar(30);not null"`
	/*
			index:idx_mobile:手机号
			unique:唯一单独的
			type:varchar(30):字符串 只能存30个字符
		    not null:不为空，必须要填写
	*/
	Role int `gorm:"column:role;default:1;type:int; comment:'1表示普通用户, 2表示管理员'"`
	/**
	gorm:"column:role(列名也就是字段名,也可以自定义其他的比如rolq，但是最好是字段名)
	default:1  默认1(普通用户)
	type:int  类型整形
	comment:'1表示普通用户, 2表示管理员' ---->赋值时要么是1 要么是2 默认是1(普通用户)

	*/
}

type Person_sd struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	//Birthday time.Time
	Age int
}

func gorm_Creating_Record() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 返回 error
	if err != nil {
		panic(err)
	}
	//自动迁移:使用表
	db.AutoMigrate(&Person_sd{})

	us := Person_sd{
		ID:   1,
		Name: "习近平",
		Age:  20,
		//Birthday: time.Now(),
	}

	updateUS := Person_sd{
		ID:   100,
		Name: "王强",
		Age:  19,
		//Birthday: time.Now(),
	}

	db.Create(&us)
	fmt.Print("id:", "姓名:", "年龄:", us.ID, us.Name, us.Age)
	//1 创建记录并更新给出的字段
	/*
		然后刷新MySQL表后会增加一行字段数据 ，把原来的Age替换为CreatedAt新字段，字段value位null
	*/
	db.Select("Name", "CreatedAt").Create(&updateUS)

	//2 创建一个记录且一同忽略传递略去的字段值
	/*比如Name不会传值
	db.Omit: 计算给要设定的Name赋值了，用Omit后也不会显示出来，会显示未null,
	*/
	user1 := Person_sd{Name: "李四", Age: 20}
	db.Omit("Name").Create(&user1)

	//静态批量插入
	/**
	因为这里只调用了Name,所以新加的3行里Age字段变量都未0
	*/
	var multipleInsert = []Person_sd{
		{Name: "名字01"},
		{Name: "名字02"},
		{Name: "名字03"},
	}
	db.Create(&multipleInsert)

	db.Clauses()

	//根据 Map 创建
	//GORM 支持根据 map[string]interface{} 和 []map[string]interface{}{} 创建记录，例如：
	//map插入一行
	db.Model(&Person_sd{}).Create(map[string]interface{}{
		"Name": "刘大爷",
		"Age":  30,
	})
	//map插入多行：
	db.Model(&Person_sd{}).Create([]map[string]interface{}{
		{"Name": "曾大爷",
			"Age": 30,
		},
		{"Name": "王大爷",
			"Age": 3,
		},
		{"Name": "黄大爷",
			"Age": 308,
		},
	})

	//标签新表
	dsn_mark := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db_mark, err_mark := gorm.Open(mysql.Open(dsn_mark), &gorm.Config{})
	// 返回 error
	if err_mark != nil {
		panic(err_mark)
	}
	//自动迁移:使用表
	db_mark.AutoMigrate(&MARKMODEL{})
	//var sub_bm = BaseModel{Id: 20}
	var mrak_model = MARKMODEL{

		OpenID: "123456",
		Role:   2,
	}

	db_mark.Create(&mrak_model)

}
