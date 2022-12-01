package grammarLearning

import (
	"encoding/json"
	"fmt"
)

//import "fmt"

type Student1 struct {
	Name     string
	Age      int
	Address1 Address1
}

type Address1 struct {
	Road     string
	Street   string
	City     string
	Province string
	Country  string
}

func MainApi() {
	Zhang3 := Student1{
		Name: "张三",
		Age:  18,
		Address1: Address1{
			Road:     "高德大厦路",
			Street:   "高普路",
			City:     "广州市",
			Province: "广东省",
			Country:  "中国(CN)",
		},
	}
	Info_of_Zhang3, err := json.Marshal(Zhang3)
	if err == nil {
		fmt.Println("接口:", string(Info_of_Zhang3))
	} else {
		fmt.Println(err)
		return
	}
}
