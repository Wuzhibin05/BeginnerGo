package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//创建结构体变量和访问结构体字段
	var person Person
	person.Name = "Alice"

	var person1 Person = Person{
		Name: "bruce",
		Age:  33,
	}
	fmt.Println(person1)
}
