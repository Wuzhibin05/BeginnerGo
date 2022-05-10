package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	// 方式一:定义直接赋值，使用值传递
	var stu1 = Student{"tom", 24}
	var stu2 = Student{
		name: "tom",
		age:  24,
	}
	stu3 := Student{"bruce", 24}
	fmt.Println(stu1, stu2, stu3)
	// 方式二: 定义赋值，使用指针
	var stu4 = &Student{"tom", 24}
	stu5 := &Student{"tom", 24}
	var stu6 *Student = &Student{
		name: "tom",
		age:  24,
	}
	stu7 := &Student{
		name: "tom",
		age:  24,
	}
	fmt.Println(*stu4, *stu5, *stu6, *stu7)
}
