package main

import "fmt"

/*
小狗案例 [学员课后练习]
1) 编写一个 Dog 结构体， 包含 name、 age、 weight 字段
2) 结构体中声明一个 say 方法， 返回 string 类型， 方法返回信息中包含所有字段值。
3) 在 main 方法中， 创建 Dog 结构体实例(变量)， 并访问 say 方法， 将调用结果打印输出。
*/
type Dog struct {
	name   string
	age    int
	weight float64
}

func (dog Dog) say() string {
	resStr := fmt.Sprintf("name=%v\nage=%v\nweight=%v", dog.name, dog.age, dog.weight)
	return resStr
}

func main() {
	var dog1 = Dog{
		name:   "tutu",
		age:    8,
		weight: 5,
	}
	res := dog1.say()
	fmt.Println(res)
}
