package main

import "fmt"

func main() {
	var monster []map[string]string
	monster = make([]map[string]string, 2, 10)
	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "蜘蛛精"
		monster[0]["age"] = "200"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["name"] = "三骨精"
		monster[1]["age"] = "400"
	}
	// 超过切片长度，越界
	//monster[2] = make(map[string]string)
	//monster[2]["name"] = "鞋子"
	//monster[2]["age"] = "500"
	//fmt.Println(monster)
	// 这里我们需要使用append函数动态鞥家切片
	newMonster := map[string]string{
		"name": "三骨精",
		"age":  "400",
	}
	monster = append(monster, newMonster)
	fmt.Println(monster)
}
