package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float64
	var isPass bool

	fmt.Println("请输入姓名 年龄，薪水，否通过考试，使用空格隔开 ")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
}
