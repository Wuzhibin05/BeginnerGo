package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//创建结构体变量和访问结构体字段
	// 方式一
	var person Person
	person.Name = "Alice"

	// 方式二
	var person1 Person = Person{
		Name: "bruce",
		Age:  33,
	}
	fmt.Println(person1.Name)
	fmt.Println(person1.Age)

	// 方式三
	var p3 *Person = new(Person)
	p3.Name = "Alice"
	p3.Age = 22
	fmt.Println(p3.Name)
	fmt.Println(p3.Age)

	// 方式四
	var p4 *Person = &Person{
		Age:  22,
		Name: "tom",
	}
	fmt.Println(p4.Name)
	fmt.Println(p4.Age)
}
