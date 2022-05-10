package main

import "fmt"

/*
1) 编写一个 Student 结构体， 包含 name、 gender、 age、 id、 score 字段， 分别为 string、string、 int、int、 float64 类型。
2) 结构体中声明一个 say 方法， 返回 string 类型， 方法返回信息中包含所有字段值。
3) 在 main 方法中， 创建 Student 结构体实例(变量)， 并访问 say 方法， 并将调用结果打印输出。
*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu Student) say() string {
	resStr := fmt.Sprintf("学生为: %v\n性别：%v\n年龄:%v\n学号：%v\n考试分数：%v", stu.name, stu.gender, stu.age, stu.id, stu.score)
	return resStr
}

func main() {
	var stu1 = Student{
		name:   "tom",
		gender: "男",
		age:    24,
		id:     20084156,
		score:  89.99,
	}
	//m1.matrix()
	res := stu1.say()
	fmt.Println(res)
}
