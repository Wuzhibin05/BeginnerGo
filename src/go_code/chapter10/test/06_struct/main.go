package main

import "fmt"
import "encoding/json"

type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	//1. 创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	//2. 将monster变量序列化为 json格式字串
	//   json.Marshal 函数中使用反射，这个讲解反射时，我会详细介绍
	fmt.Println(monster)
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误 ", err)
	}
	fmt.Println("jsonStr", string(jsonStr))

}
