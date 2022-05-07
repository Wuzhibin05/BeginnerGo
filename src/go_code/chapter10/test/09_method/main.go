package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func (s *student) String() string {
	str := fmt.Sprintf("name=%v,age=%v", s.Name, s.Age)
	return str
}

func main() {
	var stu1 = student{
		Name: "tom",
		Age:  22,
	}
	//如果你实现了 *Student 类型的 String方法，就会自动调用
	fmt.Println(&stu1)
}
