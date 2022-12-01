package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 默认只能有3个主键:只能定义3个参数
type Userinfo struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Gender   string
	Birthday time.Time
	//	Deleted  gorm.DeletedAt //如果你不想引入gorm.Model，你也可以这样启用软删除特性
	Hobby  string
	Fans   string
	Weight uint
}

func gorm1() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 返回 error
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&Userinfo{})
	u1 := Userinfo{Id: 1, Name: "张三", Gender: "男", Birthday: time.Now(), Hobby: "学习", Fans: "20", Weight: 90}
	//创建记录并更新给出的字段
	//db.Select("Name", "height", " Gender", "Hobby").Create(&Userinfo{})

	var i_data = []Userinfo{{Id: 112, Name: "李四", Gender: "女", Birthday: time.Now(), Hobby: "学习", Fans: "20", Weight: 90}, {Id: 113, Name: "王五110", Gender: "男", Birthday: time.Now(), Hobby: "打游戏", Fans: "2000", Weight: 2}}
	//db.CreateInBatches(i_data, len(i_data))
	//db.CreateInBatches(i_data, 100)
	//var i_data = []Userinfo{{Name: "张三_1"},....,{Name:"张三_1000"}}
	//db.CreateInBatches(i_data,100)

	db.Create(&u1) //创建
	db.CreateInBatches(i_data, len(i_data))
	//db.Create(&u1)
	//result.RowsAffected // 返回插入记录的条数
	for _, u1 := range i_data {
		id_value := u1.Name
		fmt.Println("name值:", id_value)

	}
	//增加一行
	db.Select("features", "Gender", "Birthday", "Hobby", "Fans", "Weight").Create(&u1)
	/*
		重点：run后记得去MySQL数据库表--对应的表userinfos表右键刷新就出来了
	*/
	//db.Select("Name").Delete(&i_data) //错误
	db.Where("name = ?", "张三").Delete(Userinfo{}) //带额外条件的删除
	/**
	Delete(&这里得写model类型名加花括号{}):
	比如:Userinfo{}
	而不能写Userinfo{}定义后返回的常量名i_data
	*/

	//不带条件的删除，只根据字段记录字段
	//db.Delete(UserModel{})

	//根据主键删除
	//db.Delete(&Userinfo{}, 11) //这个删以ID为主键值的字段记录 (单主键)
	//db.Delete(&Userinfo{}, []int{1, 11})//删除多主键(删除ID为1 和11的字段记录)

	//批量删除
	//db.Where("gender = ?", "男").Delete(Userinfo{})
	//db.Where("gender = ?", "男").Delete(Userinfo{}) //删除性别为男的记录

	//保存   -----> 保存所有字段
	//db.First(Userinfo{})
	//db.Save(Userinfo{})

	//更新多列

	db.Model(&u1).Updates(Userinfo{Name: "hello", Fans: "200粉丝"})

	// 根据 `map` 更新属性
	//db.Model(&u1).Updates(map[string]interface(){"Name":"傻逼","Hobby":"学习","Fans":"300"})
	//jl := db.First(&u1)
	fmt.Println("获取第一条记录:\n", db.Last(&u1))
	/**

	获取第一条记录（主键升序）
	db.First(&user)
	获取一条记录，没有指定排序字段
	db.Take(&user)
	  获取一条记录，没有指定排序字段
	  db.Take(&user)

	*/
	//这句放到所以操作的最后，这样就不需要手动去MySQL刷新
	db.Create(&u1)
}
