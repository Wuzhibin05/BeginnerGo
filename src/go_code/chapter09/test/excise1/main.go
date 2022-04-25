package main

import "fmt"

func main() {
	/*
		演示一个 key-value 的 value 是 map 的案例，比如： 我们要存放 3 个学生信息, 每个学生有 name 和 sex 信息。思路: map[string]map[string]string
	*/
	var studentMap = make(map[string]map[string]string)
	studentMap["stu1"] = make(map[string]string)
	studentMap["stu2"] = make(map[string]string)
	studentMap["stu3"] = make(map[string]string)

	studentMap["stu1"]["name"] = "Tom"
	studentMap["stu1"]["sex"] = "男"
	studentMap["stu1"]["address"] = "beijing"

	studentMap["stu2"]["name"] = "Bruce"
	studentMap["stu2"]["sex"] = "男"
	studentMap["stu2"]["address"] = "shanghai"

	studentMap["stu3"]["name"] = "Lily"
	studentMap["stu3"]["sex"] = "女"
	studentMap["stu3"]["address"] = "chengdu"
	fmt.Println("studentMap", studentMap)

}
