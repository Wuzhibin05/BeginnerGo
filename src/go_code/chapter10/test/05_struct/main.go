package main

import "fmt"

type Student struct {
	Name string
	Age  int
}
type Stu Student

func main() {
	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1) // 不能直接让他们相等 stu2=stu1，golang会认为他们是不同的数据类型
	fmt.Println(stu1, stu2)

}
