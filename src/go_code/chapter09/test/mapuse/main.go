package main

import "fmt"

func main() {
	// 方式一：定义后make
	var heros1 map[string]string
	heros1 = make(map[string]string)
	heros1["no1"] = "宋江"
	heros1["no2"] = "卢俊义"
	heros1["no3"] = "吴用"
	fmt.Println("heros1=", heros1)

	// 方式二：类型推导make后直再赋值
	cities := make(map[string]string)
	cities["no1"] = "成都"
	cities["no2"] = "深圳"
	cities["no3"] = "武汉"
	fmt.Println("cities=", cities)

	// 方式三：类型推导后直接赋值
	heros2 := map[string]string{
		"no1": "宋江",
		"no2": "卢俊义",
		"no3": "林冲",
	}
	heros2["no4"] = "鲁智深"
	fmt.Println("heros2=", heros2)
}
