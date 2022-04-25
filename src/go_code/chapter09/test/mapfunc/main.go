package main

import (
	"fmt"
)

func main() {
	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println("初始化map：cities=", cities)

	//因为 no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "成都"
	cities["no4"] = "拉萨"
	fmt.Println("执行修改后：cities=", cities)

	//演示删除,删除元素不存在不报错
	delete(cities, "no3")
	fmt.Println("执行删除后1：cities=", cities)
	delete(cities, "no5")
	fmt.Println("执行删除后2：cities=", cities)

	//演示map的查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("找到了no1，值为%v", val)
	} else {
		fmt.Println("未找到")
	}

	fmt.Println()
	//如果希望一次性删除所有的key
	//1. 遍历所有的key,如何逐一删除 [遍历]
	//2. 直接make一个新的空间
	for keys, _ := range cities {
		delete(cities, keys)
	}
	fmt.Println("清空后的，cities=", cities)

	cities1 := make(map[string]string)
	cities1["no1"] = "北京"
	cities1["no2"] = "天津"
	cities1["no3"] = "上海"
	fmt.Println("新的cites1=", cities1)

	cities1 = make(map[string]string)
	fmt.Println("清空后的，cities1=", cities1)

}
