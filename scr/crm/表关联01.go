package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// gorm多对多关系建立与关联模式----------------02
type Students struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Name    string
	Age     int
	Classes []Classes `gorm:"many2many:students_classes;"`
}

type Classes struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Name       string
	WidthCount int
	Princ      int
	City       string
	Students   []Students `gorm:"many2many:students_classes;"`
}

//var db *gorm.DB

// 表关联
func gorm_association_gl() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库错误", err)
		panic(err)
	}
	db.AutoMigrate(&Students{})

	//fmt.Printf("成功:%v\n", db)
	var db1 = db
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功打开数据库%v\n", db1)

	/**

	流程为建立id为1的student模型，在此模型下，连接Classes表，并与id为1的class记录创建模型。
	在student表与class表上没有任何变化，但连接表students_classes表建立了一个关系记录

	*/
	class := &Classes{Model: gorm.Model{
		ID: 1,
	},
	}

	stu := &Students{Model: gorm.Model{
		ID: 1,
	},
		Name: "学生1", Age: 100,
	}
	cv := Classes{
		ID:         1,
		Princ:      12,
		WidthCount: 1,
		Name:       "name_test",
		City:       "广州",
	}
	db.Model(&stu).Association("Classes").Append(&cv)

	db.Create(&stu)
	//os.Exit(0)
	//有了关系后，尝试通过student查找其关联的Classes
	db.Preload("Classes").First(stu,
		1)
	fmt.Println("打印表单:\n", stu) //会打印要查询类的Classes的一些字段和数据
	/*
		ID  name  WidthCount  价格
		 1 name_test 1 12

	*/

	//普通查询
	db.Model(&stu).Association("Classes").Find(&class)
	fmt.Println("普通查询:\n", cv)
	/*
		ID  name    WidthCount   价格
		 1 name_test    1            12

	*/

	//添加条件测试
	//先为stu增加一个Eng信息的记录关联
	//普通查询

	csd := &[]Classes{}
	db.Model(&stu).Association("Classes").Find(csd)
	fmt.Println("添加条件测试-普通查询:\n", csd)
	/*
		ID  name    WidthCount   价格
		 1 name_test    1            12

	*/

	//添加条件测试
	//先为stu增加一个Eng信息的记录关联
	//普通查询 --->增加条件后：

	db.Model(&stu).Where("id=2").Association("Classes").Find(csd)
	fmt.Println("添加条件测试---->增加条件后:\n", csd)
	/**
	打印:
	 []
	注解:

	由此可以发现，where语句的条件是用来限制后面的Find方法的，而不是限制前面Model的建立的，展示的实验还不足以论证结果，但是本人私下尝试，删除构造的stu中的Id，改变where中的条件，让name指向stu中的name，等方法，均证明where条件是约束后面find方法，前面model的选择，似乎只能在创建对象的时候，指定好目标
	*/

	//替换关联:
	//用一个新的关联替换当时的关联
	db.Model(&stu).Association("Classes").Replace(&Classes{Name: "替换的新名字"}) //数据库会显示新名字且加新一行数据记录
	/*
			Replace:有个弊端 那就是只替换单字段数据时，其他字段数据会根据定义的类型默认为最低的，
		   比如int的字段会 全为0
		   比如string的字段 会全为null
	*/
	//	删除关联
	db.Model(&stu).Association("Classes").Delete(&Classes{Model: gorm.Model{ID: 1}})
	//QueryMany2Many() //暂时找不到这个方法，我也不知道咋整

	//清除关联
	db.Model(&stu).Association("Classes").Clear()
	/*
		因为studnets和class关联起来而生成的students_classes表的classes_id和students_id字段的id都被clear(关联清除)

	*/
	//这里报错
	//db.Select("Age").Delete(&stu)
	//db.Select("City").Delete(&Classes{})
}
